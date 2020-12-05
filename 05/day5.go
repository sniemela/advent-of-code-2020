package main

import (
	"log"
	"math"
	"sort"

	"acme.com/aoc/utils"
)

func main() {
	boardingPasses, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	log.Printf("Part 1 = %d\n", part1(boardingPasses))
	log.Printf("Part 2 = %d\n", part2(boardingPasses))
}

func part1(boardingPasses []string) int {
	max := math.MinInt64

	for _, pass := range boardingPasses {
		id := seatID(pass)

		if id > max {
			max = id
		}
	}

	return max
}

func part2(boardingPasses []string) int {
	ids := make([]int, len(boardingPasses))

	for i, pass := range boardingPasses {
		ids[i] = seatID(pass)
	}

	sort.Ints(ids)

	for i := 1; i < len(ids); i++ {
		diff := ids[i] - ids[i-1]

		if diff > 1 {
			return ids[i] - diff + 1
		}
	}

	return 0
}

func seatID(boardingPass string) int {
	return row(boardingPass)*8 + column(boardingPass)
}

func row(boardingPass string) int {
	return decode(boardingPass[:7], 0, 127)
}

func column(boardingPass string) int {
	return decode(boardingPass[7:], 0, 7)
}

func decode(part string, lower int, upper int) int {
	for _, direction := range part[:len(part)-1] {
		switch direction {
		case 'F', 'L':
			upper = (lower + upper) / 2
		case 'B', 'R':
			lower = int(math.Ceil(float64(lower+upper) / 2))
		}
	}

	last := part[len(part)-1]

	if last == 'F' || last == 'L' {
		return lower
	}

	return upper
}
