package main

import (
	"bufio"
	"fmt"
	"os"
)

//

type coordinate struct {
	x, y int
}

type lines struct {
	b, e coordinate
}

// maxValue returns the maximum value in a list of lines
func maxValue(l []lines) int {
	max := 0
	for _, v := range l {
		if v.b.x > max {
			max = v.b.x
		}
		if v.e.x > max {
			max = v.e.x
		}
		if v.b.y > max {
			max = v.b.y
		}
		if v.e.y > max {
			max = v.e.y
		}
	}
	return max
}

// emtpyMap creates an empty map of coordinates froma list of lines
func emptyMap(l []lines) [][]int {
	max := maxValue(l)
	m := make([max][max]int, 0)
	for _, v := range l {
		for x := v.b.x; x <= v.e.x; x++ {
			for y := v.b.y; y <= v.e.y; y++ {
				m = append(m, []int{x, y})
			}
		}
	}
	return m
}

// mapCoordinates creates a matrix of coordinates from a list of lines
func mapCoordinates(l []lines) [][]int {
	m := make([][]int, 0)
	fmt.Println(m)

	for _, line := range l {
		// Horizontal line (x is the same)
		if line.b.x == line.e.x {
			for y := line.b.y; y <= line.e.y; y++ {
				m[line.b.x][y] += 1
			}
			// Vertical line (y is the same)
		} else if line.b.y == line.e.y {
			for x := line.b.x; x <= line.e.x; x++ {
				m[x][line.b.y] += 1
			}
		}
	}

	return m
}

func main() {
	s := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	fmt.Println("=== Test ===")
	c := stringsToCoordinates(s)
	fmt.Println(c)
	m := mapCoordinates(c)
	fmt.Println(m)

	fmt.Println("== Part 1 ===")
	//fmt.Println(readCoordinates("../inputs/day05"))

	fmt.Println("== Part 2 ===")
}

// readCoordinates reads a file of coordinates and returns a list of lines
func readCoordinates(filename string) []lines {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var l []lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var b, e coordinate
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &b.x, &b.y, &e.x, &e.y)
		if err != nil {
			panic(err)
		}
		l = append(l, lines{b, e})
	}

	return l
}

// stringToCoordinates converts list of strings to a list of lines
func stringsToCoordinates(s []string) []lines {
	var l []lines

	for _, line := range s {
		var b, e coordinate
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &b.x, &b.y, &e.x, &e.y)
		if err != nil {
			panic(err)
		}
		l = append(l, lines{b, e})
	}

	return l
}
