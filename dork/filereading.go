package dork

import (
	"bufio"
	"os"
	"strings"
)

// readLinesFromFile reports all the lines in the given file
func readLinesFromFile(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isEmptyLine(line) == false && strings.HasPrefix(line, "#") == false {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

// isEmptyLine reports if the given string represents an empty line
func isEmptyLine(line string) bool {
	return len(strings.TrimSpace(line)) == 0
}
