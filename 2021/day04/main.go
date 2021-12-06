package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board struct {
	rows [][]int
}

// playBingo returns the index of the winning board from a list of boards and drawing numbers
func playBingo(boards []board, drawingNumbers []int) (i int, result int) {
	for _, number := range drawingNumbers {
		for bx, b := range boards {
			for i, row := range b.rows {
				for j, cell := range row {
					if cell == number {
						b.rows[i][j] = -1

						// Obtain the values in the column (j) to check
						var column []int
						for _, r := range b.rows {
							column = append(column, r[j])
						}

						// check if the row or the Column is complete
						if isLineComplete(b.rows[i]) || isLineComplete(column) {
							//fmt.Printf("Bingo (row): %v\n", b.rows[i])
							//fmt.Printf("Bingo (col): %v\n", column)
							sumReminder := sumReminder(b)
							return bx, number * sumReminder
						}

					}

				}
			}
		}
	}
	return 0, 0
}

// isLineComplete returns true if the row is complete
// complete means all its values are negative
func isLineComplete(l []int) bool {
	for _, v := range l {
		if v != -1 {
			return false
		}
	}

	return true
}

// sumReminder returns the sum of the reminder of the board
func sumReminder(board board) int {
	var sum int

	for _, row := range board.rows {
		for _, cell := range row {
			if cell == -1 {
				continue
			}
			sum += cell
		}
	}

	return sum
}

func main() {
	drawingNumbers := []int{
		7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1,
	}

	bingo := []board{
		{
			rows: [][]int{
				{22, 13, 17, 11, 0},
				{8, 2, 23, 4, 24},
				{21, 9, 14, 16, 7},
				{6, 10, 3, 18, 5},
				{1, 12, 20, 15, 19},
			},
		},
		{
			rows: [][]int{
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6},
			},
		},
		{
			rows: [][]int{
				{14, 21, 17, 24, 4},
				{10, 16, 15, 9, 19},
				{18, 8, 23, 26, 20},
				{22, 11, 13, 6, 5},
				{2, 0, 12, 3, 7},
			},
		},
	}

	i, result := playBingo(bingo, drawingNumbers)

	fmt.Println("=== Test ===")
	fmt.Println(result)
	fmt.Println(bingo[i])

	fmt.Println("== Part 1 == <25600")
	drawingNumbers = readNumbers("../inputs/day04_numbers")
	boards := readBoards("../inputs/day04_boards_fix")
	i, result = playBingo(boards, drawingNumbers)
	fmt.Println(result)
	fmt.Println(boards[i])
	fmt.Println("== Part 2 ==")
}

// readNumbers reads a file and returns a list of numbers
func readNumbers(fileName string) []int {
	var input []int

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		for _, number := range numbers {
			n, _ := strconv.Atoi(number)
			input = append(input, n)
		}
	}

	return input
}

// readBoards reads a file and returns a list of boards
func readBoards(fileName string) []board {
	var boards []board

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows [][]int
	for scanner.Scan() {
		line := scanner.Text()

		// Read the board
		if line != "" {
			numbers := strings.Split(line, ",")
			var row []int
			for _, number := range numbers {
				n, _ := strconv.Atoi(number)
				row = append(row, n)
			}
			rows = append(rows, row)
		} else {
			boards = append(boards, board{rows: rows})
			rows = [][]int{}
		}
	}

	return boards
}
