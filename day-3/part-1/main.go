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
	file, err := os.Open("day-3/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create scanner
	scanner := bufio.NewScanner(file)

	rowsOfOne := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	rowsOfZero := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	// Read lines
	line := 0
	for scanner.Scan() {
		reading := scanner.Text()
		elements := strings.Split(reading, "")

		for i, el := range elements {
			num, _ := strconv.Atoi(el)

			if num == 0 {
				rowsOfZero[i] += 1
			} else if num == 1 {
				rowsOfOne[i] += 1
			}
		}

		line += 1
	}
	g := calcGammaBits(rowsOfOne, rowsOfZero)
	e := calcEpsilonBits(rowsOfOne, rowsOfZero)

	gm, _ := strconv.ParseInt(bString(g), 2, 64)
	em, _ := strconv.ParseInt(bString(e), 2, 64)

	fmt.Println(gm)
	fmt.Println(em)
	fmt.Println(em * gm)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func bString(i []int) (s string) {
	n := 0
	for range i {
		s += fmt.Sprint(i[n])
		n += 1
	}
	return s
}

func calcGammaBits(rowsOfOne, rowsOfZero []int) (a []int) {
	if len(rowsOfZero) != len(rowsOfOne) {
		panic("arrays not equal")
	}
	i := 0
	a = make([]int, len(rowsOfZero))
	for range rowsOfOne {
		a[i] = t(rowsOfOne[i] > rowsOfZero[i])
		i += 1
	}

	return a
}

func calcEpsilonBits(rowsOfOne, rowsOfZero []int) (a []int) {
	if len(rowsOfZero) != len(rowsOfOne) {
		panic("arrays not equal")
	}
	i := 0
	a = make([]int, len(rowsOfZero))
	for range rowsOfOne {
		a[i] = t(rowsOfOne[i] < rowsOfZero[i])
		i += 1
	}

	return a
}

func t(i bool) int {
	if i {
		return 1
	}
	return 0
}
