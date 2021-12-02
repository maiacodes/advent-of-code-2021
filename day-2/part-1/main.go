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
	depth := 0

	for scanner.Scan() {
		instruction := scanner.Text()
		parts := strings.Split(instruction, " ")
		command := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		switch command {
		case "forward":
			horizontal += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}

	fmt.Printf("H: %v\nD: %v\nR: %v\n", horizontal, depth, horizontal*depth)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
