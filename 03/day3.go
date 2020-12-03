package main

import (
	"log"

	"acme.com/aoc/utils"
)

func main() {
	world, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s", err)
	}

	log.Printf("Part 1 = %d\n", part1(world))
	log.Printf("Part 2 = %d\n", part2(world))
}

func part1(world []string) int {
	return trees(world, 3, 1)
}

func part2(world []string) int {
	return trees(world, 1, 1) * trees(world, 3, 1) * trees(world, 5, 1) * trees(world, 7, 1) * trees(world, 1, 2)
}

func trees(world []string, right int, down int) int {
	trees := 0
	x := 0

	for y := down; y < len(world); y += down {
		x = x + right
		line := world[y]
		tile := line[x%len(line)]

		if tile == '#' {
			trees++
		}
	}

	return trees
}
