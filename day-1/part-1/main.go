package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day-1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	line := 0
	puzzle := 0
	lastNo := 0

	for scanner.Scan() {
		line += 1
		measurement, _ := strconv.Atoi(scanner.Text())
		fmt.Println(fmt.Sprintf("[%v] %v", line, measurement))

		if line == 1 {
			lastNo = measurement
			continue
		}
		if lastNo < measurement {
			puzzle += 1
			fmt.Println("Increased")
		} else {
			fmt.Println("Decreased")
		}
		lastNo = measurement
	}

	fmt.Printf("--- %v increases --", puzzle)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
