package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// readFile reads a csv file and returns a list of integers
func readFile(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, n := range strings.Split(line, ",") {
			i, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, i)
		}
	}

	return numbers
}

// fuelRequired calculates the fuel required for a list of integers
func fuelRequired(numbers []int) int {
	var fuel int
	var sum int

	for _, n := range numbers {
		sum += n
	}
	// TODO: complete the challenge :cat_face:
	average, _ := math.Modf(float64(sum) / float64(len(numbers)))
	measure := fuel

	for {
		if fuel < measure {
			break
		}

	}

	return fuel
}

// fuelCost calculates the fuel cost from list of integers and a cost
func fuelCost(numbers []int, cost int) int {
	var fuel int
	var sum int

	for _, n := range numbers {
		sum += n
	}
	average, _ := math.Modf(float64(sum) / float64(len(numbers)))
	measure := fuel

	for {
		if fuel < measure {
			break
		}

	}

	return fuel
}

// main
func main() {
	testInput := []int{
		16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
	}

	fmt.Println("=== Test ===")
	fmt.Println(fuelRequired(testInput))

	fmt.Println("== Part 1 ==")
	initialState := readFile("../inputs/day07")
	fmt.Println(initialState)

	fmt.Println("== Part 2 ==")

}
