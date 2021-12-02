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
	// Open file
	file, err := os.Open("day-2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Read lines

	horizontal := 0
	aim := 0
	depth := 0

	for scanner.Scan() {
		instruction := scanner.Text()
		parts := strings.Split(instruction, " ")
		command := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		switch command {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	printStats(horizontal, depth, aim)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func printStats(horizontal, depth, aim int) {
	fmt.Printf("Horizontal: %v\nDepth: %v\nAim: %v\nResult: %v\n", horizontal, depth, aim, horizontal*depth)
}
