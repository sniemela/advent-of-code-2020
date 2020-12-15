package main

import (
	"log"
	"math"
	"strconv"

	"acme.com/aoc/utils"
)

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	log.Printf("Part 1 = %d\n", part1(lines))
	log.Printf("Part 2 = %d\n", part2(lines))
}

func part1(lines []string) int {
	ship := newShip(&point{1, 0}, &point{0, 0})

	for _, line := range lines {
		action := rune(line[0])
		input, _ := strconv.Atoi(line[1:])

		switch action {
		case 'F':
			ship.forward(input)
		case 'N', 'S', 'E', 'W':
			ship.move(action, input)
		case 'L', 'R':
			ship.turn(action, input)
		}
	}

	return ship.manhattanDistance()
}

func part2(lines []string) int {
	waypoint := &point{10, -1}
	ship := &ship{waypoint, &point{0, 0}}

	for _, line := range lines {
		action := rune(line[0])
		input, _ := strconv.Atoi(line[1:])

		switch action {
		case 'F':
			ship.forward(input)
		case 'N':
			waypoint.y -= input
		case 'S':
			waypoint.y += input
		case 'E':
			waypoint.x += input
		case 'W':
			waypoint.x -= input
		case 'L', 'R':
			waypoint.rotate(action, input)
		}
	}

	return ship.manhattanDistance()
}

type point struct {
	x int
	y int
}

func (p *point) rotate(direction rune, degrees int) {
	rotations := degrees / 90
	dx := p.x
	dy := p.y

	for i := 0; i < rotations; i++ {
		tempX := dx

		if direction == 'L' {
			dx = dy
			dy = tempX * -1
		} else if direction == 'R' {
			dx = dy * -1
			dy = tempX
		}
	}

	p.x = dx
	p.y = dy
}

type ship struct {
	direction *point
	position  *point
}

func newShip(direction *point, position *point) *ship {
	return &ship{direction, position}
}

func (s *ship) move(direction rune, units int) {
	switch direction {
	case 'N':
		s.position.y -= units
	case 'S':
		s.position.y += units
	case 'E':
		s.position.x += units
	case 'W':
		s.position.x -= units
	}
}

func (s *ship) forward(units int) {
	s.position.x += s.direction.x * units
	s.position.y += s.direction.y * units
}

func (s *ship) turn(direction rune, degrees int) {
	s.direction.rotate(direction, degrees)
}

func (s *ship) manhattanDistance() int {
	return int(math.Abs(float64(s.position.x)) + math.Abs(float64(s.position.y)))
}
