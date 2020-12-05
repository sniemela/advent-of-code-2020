package main

import (
	"log"
	"regexp"
	"strings"

	"acme.com/aoc/utils"
)

var validators = map[string]*regexp.Regexp{
	"byr": regexp.MustCompile("^19[2-9][0-9]|200[0-2]$"),
	"iyr": regexp.MustCompile("^201[0-9]|2020$"),
	"eyr": regexp.MustCompile("^202[0-9]|2030$"),
	"hgt": regexp.MustCompile("^(?:1(?:[5-8][0-9]|[9][0-3])cm)|(?:59|6[0-9]|7[0-6])in$"),
	"hcl": regexp.MustCompile("^#[a-f0-9]{6}$"),
	"ecl": regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$"),
	"pid": regexp.MustCompile("^\\d{9}$"),
	"cid": regexp.MustCompile(".*"),
}

func main() {
	lines, err := utils.ReadLines("input.txt")

	if err != nil {
		log.Fatalf("Error reading file %s", err)
	}

	var passports []map[string]string
	var buffer []string

	for _, line := range lines {
		if len(line) == 0 {
			passportLine := strings.Join(buffer, " ")

			pairs := strings.Split(passportLine, " ")
			passport := map[string]string{}

			for _, pair := range pairs {
				passport[pair[:3]] = pair[4:]
			}

			passports = append(passports, passport)
			buffer = buffer[:0]
		} else {
			buffer = append(buffer, line)
		}
	}

	log.Printf("Part 1 = %d\n", part1(passports))
	log.Printf("Part 2 = %d\n", part2(passports))
}

func part1(passports []map[string]string) int {
	validPasswords := 0

	for _, passport := range passports {
		diff := len(validators) - len(passport)
		_, cidExists := passport["cid"]

		if diff == 0 || (diff == 1 && !cidExists) {
			validPasswords++
		}
	}

	return validPasswords
}

func part2(passports []map[string]string) int {
	validPasswords := 0

	for _, passport := range passports {
		fieldsValid := true

		for key, value := range passport {
			if !validators[key].MatchString(value) {
				fieldsValid = false
				break
			}
		}

		diff := len(validators) - len(passport)
		_, cidExists := passport["cid"]

		if fieldsValid && (diff == 0 || (diff == 1 && !cidExists)) {
			validPasswords++
		}
	}

	return validPasswords
}
