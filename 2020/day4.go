package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const INPFILE = "day4.input"

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	inp := string(f)

	data := strings.Split(inp, "\n\n")

	fmt.Printf("Part 1: %s\n", part1(data))
	fmt.Printf("Part 2: %s\n", part2(data))
}

func part1(data []string) string {
	required := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	valid := 0

	for _, passport := range data {
		validFields := 0

		for _, field := range strings.FieldsFunc(passport, splitFields) {
			key := strings.Split(field, ":")[0]

			if includes(required, key) {
				validFields++
			}
		}

		if validFields == len(required) {
			valid++
		}
	}

	return fmt.Sprint(valid)
}

func part2(data []string) string {
	return "Part 2"
}

func splitFields(r rune) bool {
	return r == ' ' || r == '\n'
}

func includes(list []string, term string) bool {
	for _, l := range list {
		if term == l {
			return true
		}
	}

	return false
}
