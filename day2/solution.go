package main

import (
	"aoc2018/utils/file"
	"fmt"
)

func checksum(boxIDs []string) int {
	exactlyTwo, exactlyThree := 0, 0

	for _, id := range boxIDs {
		occurences := map[rune]int{}
		distinct := map[int]bool{}

		for _, characterRune := range id {
			occurences[characterRune]++
		}

		for _, occurence := range occurences {
			distinct[occurence] = true
		}

		if distinct[2] {
			exactlyTwo++
		}
		if distinct[3] {
			exactlyThree++
		}
	}

	return exactlyTwo * exactlyThree
}

func commonLetters(boxIDs []string) string {
	for currentIndex, currentBoxID := range boxIDs {
		for comparedIndex, comparedBoxID := range boxIDs {
			if currentIndex == comparedIndex {
				continue
			}

			common := []byte{}

			for i := 0; i < len(currentBoxID); i++ {
				if currentBoxID[i] == comparedBoxID[i] {
					common = append(common, currentBoxID[i])
				}
			}

			if len(common) == len(currentBoxID)-1 {
				return string(common[:])
			}
		}
	}
	return "no matches"
}

func main() {
	boxIDs := file.ReadLines("box-ids.txt")

	fmt.Println("part one:", checksum(boxIDs))
	fmt.Println("part two:", commonLetters(boxIDs))
}
