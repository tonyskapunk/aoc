package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	testInput := []int{
		3, 4, 3, 1, 2,
	}

	fmt.Println("=== Test ===")
	fmt.Println(lanternfishGeneration(testInput, 80))
	fmt.Println(lanternfishGenerationV2(testInput, 80))

	fmt.Println("== Part 1 ==")
	initialState := readFile("../inputs/day06")
	fmt.Println(lanternfishGeneration(initialState, 80))

	fmt.Println("== Part 2 ==")
	//fmt.Println(lanternfishGeneration(initialState, 256))
	fmt.Println(lanternfishGenerationV2(initialState, 256))
}

// lanternfishGeneration returns the number of generations for a given state
func lanternfishGeneration(state []int, generations int) int {

	for g := 0; g < generations; g++ {
		var newGen []int
		var gGen []int

		for _, n := range state {
			if n > 0 {
				gGen = append(gGen, n-1)
			} else {
				newGen = append(newGen, 8)
				gGen = append(gGen, 6)
			}
		}

		state = append(gGen, newGen...)
	}

	return len(state)
}

func lanternfishGenerationV2(state []int, generations int) int {
	generation := [9]int{}
	for _, v := range state {
		generation[v+1] += 1
	}

	for g := 0; g <= generations; g++ {
		newGeneration := [9]int{}
		newGeneration[0] = generation[1]
		newGeneration[1] = generation[2]
		newGeneration[2] = generation[3]
		newGeneration[3] = generation[4]
		newGeneration[4] = generation[5]
		newGeneration[5] = generation[6]
		newGeneration[6] = generation[7]
		newGeneration[7] = generation[8]
		// special cases
		newGeneration[6] += generation[0]
		newGeneration[8] += generation[0]

		generation = newGeneration

		//fmt.Printf("[%d] %v\n", g, generation)
	}

	var res int
	for _, v := range generation {
		res += v
	}

	return res
}

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
