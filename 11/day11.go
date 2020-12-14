package main

import (
	"log"

	"acme.com/aoc/utils"
)

type position struct {
	x int
	y int
}

var neighbourOffsets = [8]position{
	position{1, 0},
	position{-1, 0},

	position{0, -1},
	position{1, -1},
	position{-1, -1},

	position{0, 1},
	position{1, 1},
	position{-1, 1},
}

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	layout := make([][]rune, len(lines))

	for y, row := range lines {
		layout[y] = make([]rune, len(row))

		for x, pos := range row {
			if pos == 'L' {
				pos = '#'
			}
			layout[y][x] = pos
		}
	}

	log.Printf("Part 1 = %d\n", part1(layout))
	log.Printf("Part 2 = %d\n", part2(layout))
}

func part1(layout [][]rune) int {
	state := layout

	for {
		tempState := createNewState(state)
		changes := false

		for y, row := range state {
			for x, tile := range row {
				adjacent := countAdjacent(state, x, y)

				if adjacent >= 4 && tile == '#' {
					tempState[y][x] = 'L'
					changes = true
				} else if adjacent == 0 && tile == 'L' {
					tempState[y][x] = '#'
					changes = true
				}
			}
		}

		state = tempState

		if !changes {
			break
		}
	}

	return countOccupiedSeats(state)
}

func part2(layout [][]rune) int {
	state := layout

	for {
		tempState := createNewState(state)
		changes := false

		for y, row := range state {
			for x, tile := range row {
				adjacent := countAdjacent2(state, x, y)

				if adjacent >= 5 && tile == '#' {
					tempState[y][x] = 'L'
					changes = true
				} else if adjacent == 0 && tile == 'L' {
					tempState[y][x] = '#'
					changes = true
				}
			}
		}

		state = tempState

		if !changes {
			break
		}
	}

	return countOccupiedSeats(state)
}

func createNewState(layout [][]rune) [][]rune {
	newState := make([][]rune, len(layout))
	for i, row := range layout {
		newState[i] = make([]rune, len(row))
		copy(newState[i], row)
	}
	return newState
}

func countOccupiedSeats(layout [][]rune) int {
	total := 0
	for _, row := range layout {
		for _, pos := range row {
			if pos == '#' {
				total++
			}
		}
	}
	return total
}

func countAdjacent(layout [][]rune, x, y int) int {
	total := 0
	height := len(layout)
	width := len(layout[0])

	for _, pos := range neighbourOffsets {
		x0 := x + pos.x
		y0 := y + pos.y

		if x0 >= 0 && x0 < width && y0 >= 0 && y0 < height {
			if layout[y0][x0] == '#' {
				total++
			}
		}
	}

	return total
}

func countAdjacent2(layout [][]rune, x, y int) int {
	total := 0
	height := len(layout)
	width := len(layout[0])

	for _, pos := range neighbourOffsets {
		x0 := x + pos.x
		y0 := y + pos.y

		for x0 >= 0 && x0 < width && y0 >= 0 && y0 < height {
			tile := layout[y0][x0]
			if tile == '#' {
				total++
				break
			} else if tile == 'L' {
				break
			}
			x0 += pos.x
			y0 += pos.y
		}
	}

	return total
}
