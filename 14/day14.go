package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"acme.com/aoc/utils"
)

var assignmentPattern = regexp.MustCompile("(\\d+)")

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	log.Printf("Part 1 = %d\n", part1(lines))
	log.Printf("Part 2 = %d\n", part2(lines))
}

func part1(lines []string) int {
	var mask string

	memory := make(map[int]int)

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = line[len(line)-36:]
			continue
		}
		matches := assignmentPattern.FindAllStringSubmatch(line, 2)
		address, _ := strconv.Atoi(matches[0][0])
		value, _ := strconv.Atoi(matches[1][0])
		memory[address] = maskValue(mask, value, false)[0]
	}

	sum := 0
	for _, val := range memory {
		sum += val
	}

	return sum
}

func part2(lines []string) int {
	var mask string

	memory := make(map[int]int)

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = line[len(line)-36:]
			continue
		}
		matches := assignmentPattern.FindAllStringSubmatch(line, 2)
		address, _ := strconv.Atoi(matches[0][0])
		value, _ := strconv.Atoi(matches[1][0])

		for _, newAddress := range maskValue(mask, address, true) {
			memory[newAddress] = value
		}
	}

	sum := 0
	for _, val := range memory {
		sum += val
	}

	return sum
}

func maskValue(mask string, value int, floating bool) []int {
	result := make([]rune, len(mask))
	bin := binaryString(value)

	for i := 0; i < len(mask); i++ {
		m := mask[i]
		var newValue rune

		if floating {
			if m == 'X' {
				newValue = 'X'
			} else if m == '1' {
				newValue = '1'
			} else {
				newValue = rune(bin[i])
			}
		} else if m == 'X' {
			newValue = rune(bin[i])
		} else {
			newValue = rune(mask[i])
		}

		result[i] = newValue
	}

	resultString := string(result)

	if !floating {
		converted, _ := strconv.ParseInt(resultString, 2, 64)
		ret := []int{int(converted)}
		return ret
	}

	return permutations(resultString)
}

func binaryString(value int) string {
	b := strconv.FormatInt(int64(value), 2)

	bin := ""
	count := 36 - len(b)
	for i := 0; i < count; i++ {
		bin += "0"
	}
	bin += b

	return bin
}

func permutations(value string) []int {
	xCount := 0
	for i := 0; i < len(value); i++ {
		if value[i] == 'X' {
			xCount++
		}
	}

	size := int(math.Pow(float64(2), float64(xCount)))
	binStrings := make([]string, size)
	for i := 0; i < size; i++ {
		binStrings[i] = binaryString(i)
	}

	result := make([]int, size)

	for i, bin := range binStrings {
		chars := []rune(value)
		binIndex := len(bin) - xCount

		for j, c := range chars {
			if c == 'X' {
				chars[j] = rune(bin[binIndex])
				binIndex++
			}
		}

		num, _ := strconv.ParseInt(string(chars), 2, 64)
		result[i] = int(num)
	}

	return result
}
