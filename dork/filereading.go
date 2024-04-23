package dork

import (
	"bufio"
	"os"
)

// reads all the lines in the given file
func readLinesFromFile(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
