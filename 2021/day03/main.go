package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// powerConsumption calculates the power consumption of the submarine
// by reading a list of binary numbers in a string
func powerConsumption(input []string) int {
	var gammaRate string
	var epsilonRate string

	sizeOfElement := len(input[0])

	// Loop through each element of a fixed size (columns)
	for i := 0; i < sizeOfElement; i++ {
		var verticalValue []int

		// Get the element at index (column) of all input values
		// to form the verticalValue of the column
		for _, v := range input {
			e, err := strconv.Atoi(string(v[i]))
			if err != nil {
				panic(err)
			}
			verticalValue = append(verticalValue, e)
		}

		// Calculate the most and least common elements in verticalValue
		most, least := getMostLeast(verticalValue)

		gammaRate += strconv.Itoa(most)
		epsilonRate += strconv.Itoa(least)
	}

	// Transform gammaRate and epsilonRate to int from binary
	gammaRateInt, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonRateInt, _ := strconv.ParseInt(epsilonRate, 2, 64)

	//fmt.Println(gammaRateInt)
	//fmt.Println(epsilonRateInt)

	return int(gammaRateInt) * int(epsilonRateInt)
}

// getMostLeast returns the most and least common element in a list of binary numbers
func getMostLeast(input []int) (int, int) {
	var zeros int
	var ones int
	var most int
	var least int

	for _, v := range input {
		if v == 0 {
			zeros++
		} else if v == 1 {
			ones++
		}
	}

	if ones > zeros {
		most = 1
		// If we didn't count any zero then the least is 1
		if zeros == 0 {
			least = 1
		} else {
			least = 0
		}
	} else if zeros > ones {
		most = 0
		if ones == 0 {
			// If we didn't count any one then the least is 0
			least = 0
		} else {
			least = 1
		}
	} else {
		// When both zeros and ones are the same but different of 0
		// then the most is 1 and the least is 0
		most = 1
		least = 0
	}

	return most, least
}

// getOxygenGeneratoRating returns the oxygen generator rating from a list of binary numbers
func getOxygenGeneratorRating(input []string) int {
	var co2GeneratorRating int64
	var oxygenGeneratorRating int64

	oxygenLevels := input
	co2Levels := input
	sizeOfElement := len(input[0])

	// Loop through each element of a fixed size (columns)
	for i := 0; i < sizeOfElement; i++ {
		var verticalValue []int

		// Get the element at index (column) of all input values
		// to form the verticalValue of the column
		for _, v := range input {
			e, err := strconv.Atoi(string(v[i]))
			if err != nil {
				panic(err)
			}
			verticalValue = append(verticalValue, e)
		}

		// oxygenLevels
		for {
			if len(oxygenLevels) == 1 {
				break
			}
			var tmpList []string

			most, _ := getMostLeast(verticalValue)
			//fmt.Printf("o2 [%v] %v: ", len(oxygenLevels), most)

			for _, v := range oxygenLevels {
				e, _ := strconv.Atoi(string(v[i]))
				if most == int(e) {
					tmpList = append(tmpList, v)
				}
			}
			oxygenLevels = tmpList
			//fmt.Println(tmpList)
			// Don't understand why this break is needed ???
			break
		}

		// co2Levels
		for {
			if len(co2Levels) == 1 {
				break
			}
			var tmpList []string

			_, least := getMostLeast(verticalValue)
			//fmt.Printf("co2: [%v] %v: ", len(co2Levels), least)

			for _, v := range co2Levels {
				e, _ := strconv.Atoi(string(v[i]))
				if least == int(e) {
					tmpList = append(tmpList, v)
				}
			}
			co2Levels = tmpList
			//fmt.Println(tmpList)
			// Don't understand why this break is needed ???
			break
		}

		//fmt.Println(oxygenLevels)
		//fmt.Println(co2Levels)
		oxygenGeneratorRating, _ = strconv.ParseInt(oxygenLevels[0], 2, 64)
		co2GeneratorRating, _ = strconv.ParseInt(co2Levels[0], 2, 64)
		//fmt.Println(oxygenGeneratorRating)
		//fmt.Println(co2GeneratorRating)
	}

	return int(co2GeneratorRating * oxygenGeneratorRating)
}

// main is the entry point of this application
func main() {
	var i []string = []string{
		"00100",
		"11110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	// Test
	fmt.Println("== Test ==")
	fmt.Println(powerConsumption(i))
	fmt.Println(getOxygenGeneratorRating(i))

	// Part 1
	fmt.Println("== Part 1 == 3985686")
	fmt.Println(powerConsumption(readFile("../inputs/day03")))

	// Part 2
	fmt.Println("== Part 2 == ! 3990690 (! 3982715 )")
	fmt.Println(getOxygenGeneratorRating(readFile("../inputs/day03")))

}

// readFile reads a file and returns a list of strings
func readFile(fileName string) []string {
	var input []string

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
