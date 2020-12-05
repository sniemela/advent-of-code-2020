package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
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
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	content := string(bytes)
	return strings.Split(content, "\n"), nil
}
