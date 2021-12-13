package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPFILE = "day13.input"

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	inp := string(f)
	inp = strings.TrimSpace(inp)

	data := strings.Split(inp, "\n\n")
	dots := strings.Split(data[0], "\n")
	ins := strings.Split(data[1], "\n")

	fmt.Printf("Part 1: %s\n", part1(dots, ins))
	fmt.Println("Part 2:")
	part2(dots, ins)
}

func part1(dots []string, ins []string) string {
	paper := parse(dots)

	paper.Fold(ins[0])

	return fmt.Sprint(paper.Count())
}

func part2(dots []string, ins []string) {
	paper := parse(dots)

	for _, inst := range ins {
		paper.Fold(inst)
	}

	paper.Print()
}

type Solver struct {
	Paper  map[int]map[int]bool
	Height int
	Width  int
}

func parse(data []string) Solver {
	paper := make(map[int]map[int]bool)
	height := 0
	width := 0

	for _, d := range data {
		c := strings.Split(d, ",")
		c0, _ := strconv.Atoi(c[0])
		c1, _ := strconv.Atoi(c[1])
		if _, ok := paper[c1]; !ok {
			paper[c1] = make(map[int]bool)
		}
		paper[c1][c0] = true
		if height < c1 {
			height = c1
		}
		if width < c0 {
			width = c0
		}
	}

	return Solver{
		Paper:  paper,
		Height: height,
		Width:  width,
	}
}

func (s *Solver) Fold(ins string) {
	sp := strings.Split(ins, " ")
	a := strings.Split(sp[2], "=")
	direction := a[0]
	amount, _ := strconv.Atoi(a[1])

	ystart := 0
	xstart := 0

	if direction == "x" {
		xstart = amount + 1
	} else {
		ystart = amount + 1
	}

	for y := ystart; y <= s.Height; y++ {
		for x := xstart; x <= s.Width; x++ {
			if !s.Paper[y][x] {
				continue
			}

			// delete(s.Paper[y], x)
			s.Update(x, y, false)

			if direction == "y" {
				dis := (y - amount) * 2

				s.Update(x, y-dis, true)

			}

			if direction == "x" {
				dis := (x - amount) * 2

				s.Update(x-dis, y, true)
			}
		}
	}

	if direction == "x" {
		s.Width = amount
	} else {
		s.Height = amount
	}
}

func (s *Solver) Update(x, y int, v bool) {
	if _, ok := s.Paper[y]; !ok {
		s.Paper[y] = make(map[int]bool)
	}
	s.Paper[y][x] = v
}

func (s *Solver) Count() int {
	count := 0

	for _, r := range s.Paper {
		for _, c := range r {
			if c {
				count++
			}
		}
	}

	return count
}

func (s *Solver) Print() {
	for y := 0; y <= s.Height; y++ {
		for x := 0; x <= s.Width; x++ {
			if s.Paper[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
