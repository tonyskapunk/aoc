package main

import (
	"bufio"
	"fmt"
	"os"
)

type instruction struct {
	command string
	value   int
}

// getHorizontalPosition returns the horizontal position of a list of instructions
func getHorizontalPosition(instructions []instruction) int {
	var horizontalPosition int
	var depth int

	for _, instruction := range instructions {
		if instruction.command == "forward" {
			horizontalPosition += instruction.value
		} else if instruction.command == "down" {
			depth += instruction.value
		} else if instruction.command == "up" {
			depth -= instruction.value
		} else {
			fmt.Println("Unknown command:", instruction.command)
			continue
		}
	}

	return horizontalPosition * depth
}

// getHorizontalpositionWithAim returns the horizontal position of a list of instructions
func getHorizontalPositionWithAim(instructions []instruction) int {
	var horizontalPosition int
	var aim int
	var depth int

	for _, instruction := range instructions {
		if instruction.command == "forward" {
			horizontalPosition += instruction.value
			depth += aim * instruction.value
		} else if instruction.command == "down" {
			aim += instruction.value
		} else if instruction.command == "up" {
			aim -= instruction.value
		} else {
			fmt.Println("Unknown command:", instruction.command)
			continue
		}
	}

	return horizontalPosition * depth
}

// readiInstructions reads a list of instructions from a file
func readInstructions(filename string) []instruction {
	var instructions []instruction

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return instructions
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var command string
		var value int
		fmt.Sscanf(scanner.Text(), "%s %d", &command, &value)
		instructions = append(instructions, instruction{command, value})
	}

	return instructions
}

func main() {
	instructions := readInstructions("../inputs/day02")

	// Part 1
	fmt.Println("=== Part 1 ===")
	fmt.Println(getHorizontalPosition(instructions))

	// Part 2
	fmt.Println("=== Part 2 ===")
	fmt.Println(getHorizontalPositionWithAim(instructions))
}
