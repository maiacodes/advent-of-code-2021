package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(oxygenGenerator() * co2Scrubber())
}

func oxygenGenerator() int64 {
	// Open file
	file, err := os.Open("day-3/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Shit out lines into array
	readings := []string{}
	for scanner.Scan() {
		readings = append(readings, scanner.Text())
	}
	goodReadings := readings

	// Find length of binary
	columnsLength := len(strings.Split(readings[0], ""))

	fmt.Println("There are ", columnsLength, " columns")

	// Read lines
	column := 0
	iteration := 0
	for true {
		if column == columnsLength {
			break
		}

		fmt.Println("Checking column ", column)
		currentReadings := goodReadings
		goodReadings = []string{}

		common := mostCommon(currentReadings, column)

		row := 0
		rows := len(currentReadings)
		fmt.Println(currentReadings)
		for true {
			if row == rows {
				break
			}
			readingRaw := currentReadings[row]
			reading, _ := strconv.Atoi(strings.Split(readingRaw, "")[column])
			fmt.Println(reading, ": ", common == reading)
			if common == reading {
				goodReadings = append(goodReadings, readingRaw)
			}
			row += 1
		}

		column += 1
		iteration += 1
	}

	fmt.Println("---")
	fmt.Println("Oxy good no:", goodReadings[0])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res, _ := strconv.ParseInt(goodReadings[0], 2, 64)
	return res
}

func co2Scrubber() int64 {
	// Open file
	file, err := os.Open("day-3/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Shit out lines into array
	readings := []string{}
	for scanner.Scan() {
		readings = append(readings, scanner.Text())
	}
	goodReadings := readings

	// Find length of binary
	columnsLength := len(strings.Split(readings[0], ""))

	fmt.Println("There are ", columnsLength, " columns")

	// Read lines
	column := 0
	iteration := 0
	for true {
		if column == columnsLength {
			break
		}

		fmt.Println("Checking column ", column)
		currentReadings := goodReadings
		goodReadings = []string{}

		common := leastCommon(currentReadings, column)

		row := 0
		rows := len(currentReadings)
		fmt.Println(currentReadings)
		if len(currentReadings) == 1 {
			goodReadings = currentReadings
			break
		}
		for true {
			if row == rows {
				break
			}
			readingRaw := currentReadings[row]
			reading, _ := strconv.Atoi(strings.Split(readingRaw, "")[column])
			fmt.Println(reading, ": ", common == reading)
			if common == reading {
				goodReadings = append(goodReadings, readingRaw)
			}

			row += 1
		}

		column += 1
		iteration += 1
	}

	fmt.Println("---")
	fmt.Println("Scrubber good no:", goodReadings[0])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res, _ := strconv.ParseInt(goodReadings[0], 2, 64)
	return res
}

func mostCommon(readings []string, column int) int {
	row := 0
	max := len(readings)

	ones := 0
	zeros := 0

	for true {
		if max == row {
			break
		}
		readingRaw := readings[row]
		reading, err := strconv.Atoi(strings.Split(readingRaw, "")[column])
		if err != nil {
			panic(err)
		}

		if reading == 0 {
			zeros += 1
		} else if reading == 1 {
			ones += 1
		}
		row += 1
	}

	if ones >= zeros {
		return 1
	}
	return 0
}

func leastCommon(readings []string, column int) int {
	row := 0
	max := len(readings)

	ones := 0
	zeros := 0

	for true {
		if max == row {
			break
		}
		readingRaw := readings[row]
		reading, err := strconv.Atoi(strings.Split(readingRaw, "")[column])
		if err != nil {
			panic(err)
		}

		if reading == 0 {
			zeros += 1
		} else if reading == 1 {
			ones += 1
		}
		row += 1
	}

	if ones < zeros {
		return 1
	}
	return 0
}
