package main

import (
	"log"
	"regexp"
	"strconv"

	"acme.com/aoc/utils"
)

var linePattern = regexp.MustCompile("(\\d+)-(\\d+)\\s([a-z]):\\s([a-z]+)")

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s", err)
	}

	log.Printf("Part 1 = %d\n", part1(lines))
	log.Printf("Part 2 = %d\n", part2(lines))
}

func part1(lines []string) int {
	validPasswords := 0

	for _, line := range lines {
		if groups := linePattern.FindStringSubmatch(line); groups != nil {
			min, _ := strconv.Atoi(groups[1])
			max, _ := strconv.Atoi(groups[2])
			letter := groups[3]
			password := groups[4]

			letterCount := 0

			for i := 0; i < len(password); i++ {
				if string(password[i]) == letter {
					letterCount++
				}
			}

			if letterCount >= min && letterCount <= max {
				validPasswords++
			}
		}
	}

	return validPasswords
}

func part2(lines []string) int {
	validPasswords := 0

	for _, line := range lines {
		if groups := linePattern.FindStringSubmatch(line); groups != nil {
			x0, _ := strconv.Atoi(groups[1])
			x1, _ := strconv.Atoi(groups[2])
			letter := groups[3]
			password := groups[4]

			if (string(password[x0-1]) == letter) != (string(password[x1-1]) == letter) {
				validPasswords++
			}
		}
	}

	return validPasswords
}
