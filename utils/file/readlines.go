package file

import (
	"bufio"
	"os"
)

// ReadLines parses file contents
func ReadLines(filename string) (lines []string) {
	file, _ := os.Open(filename)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}
