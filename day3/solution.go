package main

import (
	"aoc2018/utils/file"
	"fmt"
	"regexp"
	"strconv"
)

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func newClaim(input string) claim {
	pattern := regexp.MustCompile("#(.*) @ (.*),(.*): (.*)x(.*)")
	matches := pattern.FindStringSubmatch(input)

	id, _ := strconv.Atoi(matches[1])
	x, _ := strconv.Atoi(matches[2])
	y, _ := strconv.Atoi(matches[3])
	width, _ := strconv.Atoi(matches[4])
	height, _ := strconv.Atoi(matches[5])

	return claim{id, x, y, width, height}
}

func (c claim) points() []string {
	points := []string{}

	for x := c.x; x < c.x+c.width; x++ {
		for y := c.y; y < c.y+c.height; y++ {
			points = append(points, strconv.Itoa(x)+"_"+strconv.Itoa(y))
		}
	}

	return points
}

func squareInchOverlap(claims []claim) (overlap int) {
	claimed := getClaimedAreas(claims)

	for _, amount := range claimed {
		if amount > 1 {
			overlap++
		}
	}

	return
}

func uniqueClaim(claims []claim) int {
	claimed := getClaimedAreas(claims)

	for _, claim := range claims {
		unique := true

		for _, point := range claim.points() {
			if claimed[point] > 1 {
				unique = false
			}
		}

		if unique {
			return claim.id
		}
	}

	return -1
}

func getClaimedAreas(claims []claim) map[string]int {
	claimed := map[string]int{}

	for _, claim := range claims {
		for _, point := range claim.points() {
			claimed[point]++
		}
	}

	return claimed
}

func getClaims(lines []string) []claim {
	claims := []claim{}

	for _, line := range lines {
		claims = append(claims, newClaim(line))
	}

	return claims
}

func main() {
	claims := getClaims(file.ReadLines("claims.txt"))

	fmt.Println("part one:", squareInchOverlap(claims))
	fmt.Println("part two:", uniqueClaim(claims))
}
