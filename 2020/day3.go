package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const INPFILE = "day3.input"

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	inp := string(f)

	data := strings.Split(inp, "\n")
	data = data[:len(data)-1]

	fmt.Printf("Part 1: %s\n", part1(data))
	fmt.Printf("Part 2: %s\n", part2(data))
}

func part1(data []string) string {
	x := 0
	trees := 0

	for y := range data {
		if data[y][x] == '#' {
			trees++
		}

		x = (x + 3) % len(data[0])
	}

	return fmt.Sprint(trees)
}

func part2(data []string) string {
	return "Part 2"
}
