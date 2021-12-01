package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open file
	file, err := os.Open("day-1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Shit out lines into array
	readings := []int{}
	for scanner.Scan() {
		//line += 1
		measurement, _ := strconv.Atoi(scanner.Text())
		readings = append(readings, measurement)
	}

	increment := 0
	puzzle := 0
	lastSum := 0
	slidingWindow := [3]int{0, 0, 0}

	// Start countin'
	for true {
		// Kill when at end of da list
		if len(readings) == increment+2 {
			break
		}

		// Make the sliding window
		slidingWindow = [3]int{readings[increment], readings[increment+1], readings[increment+2]}

		// Make sum and increment counter
		sum := slidingWindow[0] + slidingWindow[1] + slidingWindow[2]
		fmt.Printf("[%v, %v, %v] = %v\n", slidingWindow[0], slidingWindow[1], slidingWindow[2], sum)

		// Check if this sum larger than last
		if lastSum != 0 && sum > lastSum {
			puzzle += 1
		}

		// Go again lol
		lastSum = sum
		increment += 1
	}

	fmt.Printf("--- %v increases --", puzzle)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
