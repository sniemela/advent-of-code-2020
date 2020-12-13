package main

import (
	"log"
	"sort"

	"acme.com/aoc/utils"
)

const preamble = 25

func main() {
	numbers, err := utils.ReadInts("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	log.Printf("Part 1 = %d\n", part1(numbers))
	log.Printf("Part 2 = %d\n", part2(numbers))
}

func part1(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		slice := numbers[i : i+preamble]
		sum := numbers[i+preamble]
		slice = slice[:preamble]

		set := make(map[int]int)
		for _, n := range slice {
			set[n] = n
		}

		valid := false

		for j := 0; j < len(slice); j++ {
			adder := sum - slice[j]

			if _, exists := set[adder]; exists {
				valid = true
				break
			}
		}

		if !valid {
			return sum
		}
	}

	return 0
}

func part2(numbers []int) int {
	invalid := part1(numbers)

	startIndex := 0
	endIndex := 0
	sum := 0
	i := 0

	for i < len(numbers) {
		sum += numbers[i]

		if sum == invalid {
			endIndex = i
			break
		} else if sum > invalid {
			startIndex++
			i = startIndex
			sum = 0
			continue
		}

		i++
	}

	resultSet := numbers[startIndex : endIndex+1]
	sort.Ints(resultSet)

	return resultSet[0] + resultSet[len(resultSet)-1]
}
