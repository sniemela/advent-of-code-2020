package main

import (
	"log"
	"sort"

	"acme.com/aoc/utils"
)

func main() {
	expenses, err := utils.ReadInts("input.txt")

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	sort.Ints(expenses)

	log.Printf("Part 1 = %d\n", part1(expenses))
	log.Printf("Part 2 = %d\n", part2(expenses))
}

func part1(expenses []int) int {
	for i := 0; i < len(expenses); i++ {
		index := searchYearIndex(expenses, expenses[i])

		if index != -1 {
			return expenses[i] * expenses[index]
		}
	}
	return 0
}

func part2(expenses []int) int {
	for i := 0; i < len(expenses); i++ {
		for j := i; j < len(expenses); j++ {
			index := searchYearIndex(expenses, expenses[i]+expenses[j])

			if index != -1 {
				return expenses[i] * expenses[j] * expenses[index]
			}
		}
	}
	return 0
}

func searchYearIndex(expenses []int, expense int) int {
	l := 0
	r := len(expenses) - 1

	for l <= r {
		m := (l + r) / 2
		sum := expenses[m] + expense

		if sum < 2020 {
			l = m + 1
		} else if sum > 2020 {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}
