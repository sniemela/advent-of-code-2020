package utils

import (
	"bufio"
	"os"
	"strconv"
)

// ReadInts reads all integers from the file
func ReadInts(file string) ([]int, error) {
	lines, err := ReadLines(file)
	if err != nil {
		return nil, err
	}

	ints := make([]int, len(lines))

	for i, line := range lines {
		converted, err := strconv.Atoi(line)

		if err != nil {
			return nil, err
		}

		ints[i] = converted
	}

	return ints, nil
}

// ReadLines reads all lines from the file
func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
