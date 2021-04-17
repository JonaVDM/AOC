package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPFILE = "day6.input"

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
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, inst := range data {
		parts := strings.Split(inst, " ")

		if parts[0] == "turn" {
			parts = parts[1:]
		}

		xy1 := strings.Split(parts[1], ",")
		x1, _ := strconv.Atoi(xy1[0])
		y1, _ := strconv.Atoi(xy1[1])

		xy2 := strings.Split(parts[3], ",")
		x2, _ := strconv.Atoi(xy2[0])
		y2, _ := strconv.Atoi(xy2[1])

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				if parts[0] == "toggle" {
					grid[x][y] = !grid[x][y]
				}

				if parts[0] == "on" {
					grid[x][y] = true
				}

				if parts[0] == "off" {
					grid[x][y] = false
				}
			}
		}
	}

	on := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] {
				on++
			}
		}
	}

	return fmt.Sprint(on)
}

func part2(data []string) string {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, inst := range data {
		parts := strings.Split(inst, " ")

		if parts[0] == "turn" {
			parts = parts[1:]
		}

		xy1 := strings.Split(parts[1], ",")
		x1, _ := strconv.Atoi(xy1[0])
		y1, _ := strconv.Atoi(xy1[1])

		xy2 := strings.Split(parts[3], ",")
		x2, _ := strconv.Atoi(xy2[0])
		y2, _ := strconv.Atoi(xy2[1])

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				if parts[0] == "toggle" {
					grid[x][y] += 2
				}

				if parts[0] == "on" {
					grid[x][y] += 1
				}

				if parts[0] == "off" {
					if grid[x][y] > 0 {
						grid[x][y] -= 1
					}
				}
			}
		}
	}

	on := 0
	for x := range grid {
		for y := range grid[x] {
			on += grid[x][y]
		}
	}

	return fmt.Sprint(on)
}
