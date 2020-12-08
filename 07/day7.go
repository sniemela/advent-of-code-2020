package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"acme.com/aoc/utils"
)

var bagPattern = regexp.MustCompile("(\\d+)([^.,]+)")

type bag struct {
	count int
	name  string
}

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	bags := make(map[string][]bag)

	for _, line := range lines {
		parts := strings.Split(line, "contain")
		name := bagName(parts[0])

		if matches := bagPattern.FindAllStringSubmatch(parts[1], -1); matches != nil {
			for _, groups := range matches {
				n, _ := strconv.Atoi(groups[1])
				innerBagName := bagName(groups[2])
				content := bags[name]
				bags[name] = append(content, bag{n, innerBagName})
			}
		}
	}

	log.Println(part1(bags))
	log.Println(part2(bags))
}

func bagName(fullName string) string {
	bag := strings.TrimSpace(fullName)
	return bag[:strings.LastIndex(bag, " ")]
}

func part1(bags map[string][]bag) int {
	queue := []string{"shiny gold"}
	outermost := make(map[string]bool)

	for len(queue) > 0 {
		bagToCheck := queue[0]
		queue = queue[1:]
		for bag, content := range bags {
			for _, innerBag := range content {
				if innerBag.name == bagToCheck && innerBag.count > 0 {
					queue = append(queue, bag)
					outermost[bag] = true
					break
				}
			}
		}
	}

	return len(outermost)
}

func part2(bags map[string][]bag) int {
	return totalBags(bags, "shiny gold") - 1
}

func totalBags(bags map[string][]bag, name string) int {
	total := 1
	for _, bag := range bags[name] {
		total += bag.count * totalBags(bags, bag.name)
	}
	return total
}
