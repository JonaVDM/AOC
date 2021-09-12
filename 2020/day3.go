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
	return fmt.Sprint(run(3, 1, data))
}

func part2(data []string) string {
	slopes := []struct{ r, d int }{
		{r: 1, d: 1},
		{r: 3, d: 1},
		{r: 5, d: 1},
		{r: 7, d: 1},
		{r: 1, d: 2},
	}

	total := 1

	for _, slope := range slopes {
		trees := run(slope.r, slope.d, data)
		log.Print(trees)
		total *= trees
	}

	return fmt.Sprint(total)
}

func run(right int, down int, hill []string) int {
	x := 0
	trees := 0

	for y := 0; y < len(hill); y += down {
		if hill[y][x] == '#' {
			trees++
		}

		x = (x + right) % len(hill[0])
	}

	return trees
}
