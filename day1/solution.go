package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(filename string) (lines []string) {
	file, _ := os.Open(filename)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}

func resultingFrequency(changes []string) (result int) {
	for _, change := range changes {
		value, _ := strconv.Atoi(change)
		result += value
	}
	return
}

func frequencyReachedTwice(changes []string) (result int) {
	frequencies := map[int]bool{0: true}

	for {
		for _, change := range changes {
			value, _ := strconv.Atoi(change)
			result += value

			if frequencies[result] {
				return
			}

			frequencies[result] = true
		}
	}
}

func main() {
	changes := readLines("frequency-changes.txt")

	fmt.Println("part one:", resultingFrequency(changes))
	fmt.Println("part two:", frequencyReachedTwice(changes))
}
