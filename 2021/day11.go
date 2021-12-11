package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPFILE = "day11.input"

// const INPFILE = "sample"

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
	grid := parse(data)

	for i := 0; i < 100; i++ {
		grid.Tick()
	}

	return fmt.Sprint(grid.GlowCount)
}

func part2(data []string) string {
	grid := parse(data)
	stepCounter := 0

	for !grid.Bright() {
		stepCounter++
		grid.Tick()
	}

	return fmt.Sprint(stepCounter)
}

func parse(data []string) Solver {
	grid := make([][]int, 0)
	for _, r := range data {
		row := make([]int, 0)
		for _, s := range strings.Split(r, "") {
			sq, _ := strconv.Atoi(s)
			row = append(row, sq)
		}
		grid = append(grid, row)
	}

	return Solver{
		Data:      data,
		Grid:      grid,
		GlowCount: 0,
	}
}

type Solver struct {
	Data      []string
	Grid      [][]int
	HasGlowed map[int]map[int]bool
	GlowCount int
}

func (s *Solver) CheckNine(x, y int) {
	if _, ok := s.HasGlowed[y]; !ok {
		s.HasGlowed[y] = make(map[int]bool)
	}
	if s.HasGlowed[y][x] {
		return
	}

	s.Grid[y][x] += 1

	if s.Grid[y][x] <= 9 {
		return
	}

	s.HasGlowed[y][x] = true

	for r := y - 1; r <= y+1; r++ {
		for c := x - 1; c <= x+1; c++ {
			if r >= 0 && r < len(s.Grid) && c >= 0 && c < len(s.Grid[0]) {
				s.CheckNine(c, r)
			}
		}
	}
}

func (s *Solver) Tick() {
	for y, row := range s.Grid {
		for x := range row {
			s.Grid[y][x] += 1
		}
	}

	s.HasGlowed = make(map[int]map[int]bool)

	for y, row := range s.Grid {
		for x := range row {
			if s.Grid[y][x] > 9 {
				s.CheckNine(x, y)
			}
		}
	}

	for y, row := range s.Grid {
		for x := range row {
			if s.Grid[y][x] > 9 {
				s.GlowCount += 1
				s.Grid[y][x] = 0
			}
		}
	}
}

func (s *Solver) Print() {
	for _, row := range s.Grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Print("\n")
	}

	fmt.Println()
}

func (s *Solver) Bright() bool {
	for _, row := range s.Grid {
		for _, col := range row {
			if col != 0 {
				return false
			}
		}
	}

	return true
}
