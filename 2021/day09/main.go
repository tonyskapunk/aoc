package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// main
func main() {
	testInput := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}

	fmt.Println("=== Test ===")
	l := findLowest(stringToInt(testInput))
	fmt.Println(riskLevel(l))

	fmt.Println("== Part 1 == ")
	lowest := findLowest(readFile("../inputs/day09"))
	//fmt.Println(lowest)
	fmt.Println(riskLevel(lowest))

	fmt.Println("== Part 2 ==")
}

// readFile reads a file and returns its content as list of integers
func readFile(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var l [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var n []int
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			e, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			n = append(n, e)
		}
		l = append(l, n)
	}

	return l
}

// stringToInt converts a list of strings to a list of integers
func stringToInt(input []string) [][]int {
	var l [][]int
	for _, line := range input {
		var n []int
		for i := 0; i < len(line); i++ {
			e, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			n = append(n, e)
		}
		l = append(l, n)
	}

	return l
}

// findLowest returns the lowest numbers in a matrix of integers
func findLowest(l [][]int) []int {
	var lowest []int
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l[i]); j++ {
			var low = -1
			var bottom = 1
			var top = 1
			var left = 1
			var right = 1
			if j == 0 {
				left = 0
			}
			if i == 0 {
				top = 0
			}
			if j == len(l[i])-1 {
				right = 0
			}
			if i == len(l)-1 {
				bottom = 0
			}

			if left == 0 {
				// Only compare to the right
				// top and bottom are 0
				if top == 0 {
					if l[i][j] < l[i][j+right] && l[i][j] < l[i+bottom][j] {
						low = l[i][j]
					}
				}
				if bottom == 0 {
					if l[i][j] < l[i][j+right] && l[i][j] < l[i-top][j] {
						low = l[i][j]
					}
				} else {
					if l[i][j] < l[i][j+right] && l[i][j] < l[i-top][j] && l[i][j] < l[i+bottom][j] {
						low = l[i][j]
					}
				}
			} else if right == 0 {
				// Only compare to the left
				// top and bottom are 0
				if top == 0 {
					if l[i][j] < l[i][j-left] && l[i][j] < l[i+bottom][j] {
						low = l[i][j]
					}
				}
				if bottom == 0 {
					if l[i][j] < l[i][j-left] && l[i][j] < l[i-top][j] {
						low = l[i][j]
					}
				} else {
					if l[i][j] < l[i][j-left] && l[i][j] < l[i-top][j] && l[i][j] < l[i+bottom][j] {
						low = l[i][j]
					}
				}
			} else if top == 0 {
				// Only compare to the bottom, not the corners(left or right)
				if l[i][j] < l[i][j+right] && l[i][j] < l[i][j-left] && l[i][j] < l[i+bottom][j] {
					low = l[i][j]
				}
			} else if bottom == 0 {
				// Only compare to the top, not the corners(left or right)
				if l[i][j] < l[i][j+right] && l[i][j] < l[i][j-left] && l[i][j] < l[i-top][j] {
					low = l[i][j]
				}
			} else {
				if l[i][j] < l[i][j+right] && l[i][j] < l[i][j-left] && l[i][j] < l[i+bottom][j] && l[i][j] < l[i-top][j] {
					low = l[i][j]
				}
			}

			if low != -1 {
				//fmt.Printf("%d<", low)
				lowest = append(lowest, low)
			} else {
				//fmt.Printf("%d ", l[i][j])
			}
		}
		//fmt.Println()
	}
	return lowest
}

// riskLevel returns the risk level of a list of integers
func riskLevel(l []int) int {
	var sum int
	for _, e := range l {
		sum += e
	}
	return sum + len(l)
}
