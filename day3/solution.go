package main

import (
	"aoc2018/utils/file"
	"fmt"
	"regexp"
	"strconv"
)

type claim struct {
	id     string
	x      int
	y      int
	width  int
	height int
}

func newClaim(input string) *claim {
	pattern := regexp.MustCompile("#(.*) @ (.*),(.*): (.*)x(.*)")
	matches := pattern.FindStringSubmatch(input)

	id := matches[1]
	x, _ := strconv.Atoi(matches[2])
	y, _ := strconv.Atoi(matches[3])
	width, _ := strconv.Atoi(matches[4])
	height, _ := strconv.Atoi(matches[5])

	return &claim{id, x, y, width, height}
}

func squareInchOverlap(lines []string) (overlap int) {
	claimedAreas := map[int]map[int]int{}

	for _, line := range lines {
		claim := newClaim(line)

		for x := claim.x; x < claim.x+claim.width; x++ {
			if claimedAreas[x] == nil {
				claimedAreas[x] = map[int]int{}
			}

			for y := claim.y; y < claim.y+claim.height; y++ {
				if claimedAreas[x][y] == 1 {
					overlap++
				}

				claimedAreas[x][y]++
			}
		}
	}

	return
}

func main() {
	claims := file.ReadLines("claims.txt")

	fmt.Println("part one:", squareInchOverlap(claims))
}
