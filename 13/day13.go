package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	"acme.com/aoc/utils"
)

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	earliest, _ := strconv.Atoi(lines[0])
	ids := strings.Split(lines[1], ",")

	log.Printf("Part 1 = %d\n", part1(earliest, ids))
	log.Printf("Part 2 = %d\n", part2(ids))
}

func part1(earliest int, ids []string) int {
	bus := 0
	min := math.MaxInt64

	for _, busID := range ids {
		if busID == "x" {
			continue
		}

		id, _ := strconv.Atoi(busID)
		mod := earliest % id
		waitingTime := earliest + id - mod

		if waitingTime < min {
			min = waitingTime
			bus = id
		}
	}

	return (min - earliest) * bus
}

func part2(ids []string) int {
	n, _ := strconv.Atoi(ids[0])
	timestamp := 0

	for i := 1; i < len(ids); i++ {
		id := ids[i]

		if id != "x" {
			busID, _ := strconv.Atoi(id)

			// Chinese reminder theorem
			for (timestamp+i)%busID != 0 {
				timestamp += n
			}

			n *= busID
		}
	}

	return timestamp
}
