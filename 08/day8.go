package main

import (
	"log"
	"strconv"

	"acme.com/aoc/utils"
)

type instruction struct {
	op  string
	arg int
}

type console struct {
	bootCode []*instruction
	result   int
}

func newConsole(instructions []*instruction) *console {
	return &console{instructions, 0}
}

func (c *console) boot() bool {
	executed := make([]int, len(c.bootCode))
	ip := 0
	c.result = 0

	for ip < len(c.bootCode) && executed[ip] < 1 {
		inst := c.bootCode[ip]
		executed[ip]++

		switch inst.op {
		case "acc":
			c.result += inst.arg
			ip++
		case "jmp":
			ip += inst.arg
		case "nop":
			ip++
		}
	}

	return ip == len(c.bootCode)
}

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	instructions := make([]*instruction, len(lines))
	for i, line := range lines {
		op, arg := parseInstruction(line)
		instructions[i] = &instruction{op, arg}
	}

	gameConsole := newConsole(instructions)

	log.Printf("Part 1 = %d\n", part1(gameConsole))
	log.Printf("Part 2 = %d\n", part2(gameConsole))
}

func parseInstruction(instruction string) (string, int) {
	operation := instruction[:3]
	argument, _ := strconv.Atoi(instruction[5:len(instruction)])

	if instruction[4] == '-' {
		argument *= -1
	}

	return operation, argument
}

func part1(c *console) int {
	c.boot()
	return c.result
}

func part2(c *console) int {
	i := 0
	prevOp := ""

	for !c.boot() {
		if prevOp != "" {
			c.bootCode[i-1].op = prevOp
		}

		prevOp = ""

		for i < len(c.bootCode) {
			inst := c.bootCode[i]
			newOp := ""

			if inst.arg != 0 {
				if inst.op == "jmp" {
					prevOp = "jmp"
					newOp = "nop"
				} else if inst.op == "nop" {
					prevOp = "nop"
					newOp = "jmp"
				}
			}

			i++

			if newOp != "" {
				inst.op = newOp
				break
			}
		}
	}

	return c.result
}
