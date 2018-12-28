package main

import (
	"aoc2018/utils/file"
	"fmt"
	"strconv"
)

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
	changes := file.ReadLines("frequency-changes.txt")

	fmt.Println("part one:", resultingFrequency(changes))
	fmt.Println("part two:", frequencyReachedTwice(changes))
}
