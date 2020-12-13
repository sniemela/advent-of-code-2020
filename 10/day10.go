package main

import (
	"log"
	"sort"

	"acme.com/aoc/utils"
)

func main() {
	adapters, err := utils.ReadInts("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	adapters = append(adapters, 0)
	copy(adapters[1:], adapters)
	adapters[0] = 0

	log.Printf("Part 1 = %d\n", part1(adapters))
	log.Printf("Part 2 = %d\n", part2(adapters))
}

func part1(adapters []int) int {
	return differences(adapters, 1) * differences(adapters, 3)
}

func differences(adapters []int, jolt int) int {
	currentJolt := 0
	differences := 0

	for i := 0; i < len(adapters); i++ {
		diff := adapters[i] - currentJolt

		if diff <= 3 {
			currentJolt = adapters[i]

			if diff == jolt {
				differences++
			}
		}
	}

	return differences
}

func part2(adapters []int) int {
	result := make([]int, adapters[len(adapters)-1]+1)
	result[0] = 1

	for _, i := range adapters[1:] {
		result[i] = result[resolveIndex(result, i-3)] + result[resolveIndex(result, i-2)] + result[resolveIndex(result, i-1)]
	}

	return result[len(result)-1]
}

func resolveIndex(result []int, index int) int {
	if index < 0 {
		return len(result) + index
	}
	return index
}
