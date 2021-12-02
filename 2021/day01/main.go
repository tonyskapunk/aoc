package main

import (
	"bufio"
	"fmt"
	"os"
)

// countIncrements returns the number of increments of previous values in a list
func countIncrements(list []int) int {
	count := 0
	for i := 1; i < len(list); i++ {
		if list[i] > list[i-1] {
			count++
		}
	}
	return count
}

// listOfSums returns a list of the sums of three consecutive numbers in a list
func listOfSums(list []int) []int {
	var sums []int
	for i := 0; i < len(list)-2; i++ {
		sums = append(sums, list[i]+list[i+1]+list[i+2])
	}
	return sums
}

func main() {
	f, err := os.Open("../inputs/day01")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	var list []int
	lines := bufio.NewScanner(f)
	for lines.Scan() {
		var value int
		_, err := fmt.Sscanf(lines.Text(), "%d", &value)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		list = append(list, value)
	}

	// Part 1
	fmt.Println("=== Part 1 ===")
	fmt.Println(countIncrements(list))

	// Part 2
	fmt.Println("=== Part 2 ===")
	list2 := listOfSums(list)
	fmt.Println(countIncrements(list2))

}
