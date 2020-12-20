package main

import (
	"log"
)

func main() {
	log.Printf("Part 1 = %d\n", play([]int{0, 3, 6}, 2020))
	log.Printf("Part 2 = %d\n", play([]int{18, 8, 0, 5, 4, 1, 20}, 30000000))
}

func play(initial []int, stopRound int) int {
	memory := make([]int, stopRound)
	turn := 1

	for _, number := range initial[:len(initial)-1] {
		memory[number] = turn
		turn++
	}

	var prev int
	current := initial[len(initial)-1]
	stopRound = stopRound + 1

	for turn != stopRound {
		prev = current

		if prevTurn := memory[current]; prevTurn != 0 {
			current = turn - prevTurn
		} else {
			current = 0
		}

		memory[prev] = turn
		turn++
	}

	return prev
}
