package main

import (
	"log"

	"acme.com/aoc/utils"
)

func main() {
	answers, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	log.Printf("Part 1 = %d\n", part1(answers))
	log.Printf("Part 2 = %d\n", part2(answers))
}

func part1(questions []string) int {
	allAnswers := map[rune]int{}
	sum := 0

	for _, question := range questions {
		if len(question) == 0 {
			allAnswers = map[rune]int{}
		} else {
			for _, q := range question {
				if _, exists := allAnswers[q]; !exists {
					allAnswers[q] = 1
					sum++
				}
			}
		}
	}

	return sum
}

func part2(questions []string) int {
	groupSize := 0
	allAnswers := map[rune]int{}
	sum := 0

	for i, question := range questions {
		if len(question) == 0 || i+1 == len(questions) {
			for _, answers := range allAnswers {
				sum += answers / groupSize
			}
			allAnswers = map[rune]int{}
			groupSize = 0
		} else {
			for _, q := range question {
				allAnswers[q]++
			}
			groupSize++
		}
	}

	return sum
}
