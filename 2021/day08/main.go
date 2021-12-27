package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type signal struct {
	digits []string
	output []string
}

func main() {
	testInput := []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}

	fmt.Println("=== Test ===")
	t := parseInput(testInput)
	fmt.Println(len(guessOutput(t)))

	fmt.Println("== Part 1 ===")
	t = parseInput(readFile("../inputs/day08"))
	fmt.Println(t)
	fmt.Println(len(guessOutput(t)))

	fmt.Println("== Part 2 ===")
}

// readFile reads a file with a list of strings and returns a slice of strings
func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// parseInput parses a slice of strings into an individually sorted slice of signals
func parseInput(input []string) []signal {
	var parsedInput []signal

	for _, line := range input {
		var s signal
		parts := strings.Split(line, " | ")
		s.digits = strSort(strings.Split(parts[0], " "))
		s.output = strSort(strings.Split(parts[1], " "))

		parsedInput = append(parsedInput, s)
	}

	return parsedInput
}

/// strSort sorts a slice of strings
func strSort(s []string) []string {
	var sorted []string

	for _, v := range s {
		sorted = append(sorted, sortString(v))
	}

	return sorted
}

// sortString sorts a string by its runes
func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))

	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

//

// guessOutput guesses the output digits of a list of signals
func guessOutput(signals []signal) []int {
	one := 2
	eight := 7
	four := 4
	seven := 3
	var digits []int

	for _, s := range signals {
		for _, d := range s.output {
			n := -1
			if len(d) == one {
				n = 1
			} else if len(d) == eight {
				n = 8
			} else if len(d) == four {
				n = 4
			} else if len(d) == seven {
				n = 7
			}
			if n != -1 {
				digits = append(digits, n)
			}
		}
	}

	return digits
}

// easyMapping maps a signal to a number
func easyMapping(s signal) [10]string {
	m := [10]string{}

	for _, d := range s.digits {
		if len(d) == 2 {
			// One
			m[1] = d
		} else if len(d) == 3 {
			// Seven
			m[7] = d
		} else if len(d) == 4 {
			// Four
			m[4] = d
		} else if len(d) == 7 {
			// Eight
			m[8] = d
		}
	}
	return m
}

// mapThree maps the number three once easyMapping is done
func mapThree(s signal, m [10]string) {
	for _, d := range s.digits {
		if strings.Contains(d, m[1]) {
			m[3] = d
		}
	}
}

// mapNine maps the number nine once easyMapping is done
func mapNine(s signal, m [10]string) string {
	var topLeft string
	for _, d := range s.digits {
		if strings.Contains(d, m[3]) {
			m[9] = d

			for _, c := range strings.Split(m[9], "") {
				if !strings.Contains(c, m[3]) {
					topLeft = c
				}
			}
		}
	}
	return topLeft
}

// mapTwoFive maps a signal to a number
func mapTwoFive(s signal, c string, m [10]string) {
	for _, d := range s.digits {
		if len(d) == 5 {
			if strings.Contains(d, c) {
				m[5] = d
			} else {
				m[2] = d
			}
		}
	}
}

// mapSixZero maps a signal to a number
func mapSixZero(s signal, m [10]string) {
	for _, d := range s.digits {
		if len(d) == 6 {
			if strings.Contains(d, m[5]) {
				m[6] = d
			} else {
				m[0] = d
			}
		}
	}
}

// complexMapping
func decodeDigits(s signal) [10]string {
	m := easyMapping(s)

	mapThree(s, m)
	topLeft := mapNine(s, m)
	mapTwoFive(s, topLeft, m)
	mapSixZero(s, m)

	return m
}

/*
 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc

So, the unique signal patterns would correspond to the following digits:

    ab: 1      # these two are in 3
    abef: 4
    abd: 7
    abcdefg: 8

    abcdf: 3  # if 1 is here then we know is a 3
    abcdef: 9 # if 3 is here then we know is a 9, the diff is top-left

	# when len is 5:
    bcdfe: 5  # if contains top-left then is a 5
    acdfg: 2  # else is 2

	# when len is 6:
    bcdefg: 6 # if 5 is here then we know is 6
    abcdeg: 0 # else is 0

*/
