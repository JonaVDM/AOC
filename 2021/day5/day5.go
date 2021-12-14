package day5

import (
	"fmt"

	"github.com/jonavdm/aoc/lib"
)

func (s *Solver) PartOne() {
	filtered := make([]Line, 0)

	for _, l := range s.lines {
		if l.x1 == l.x2 || l.y1 == l.y2 {
			filtered = append(filtered, l)
		}
	}

	fmt.Println("Part one: ", solve(filtered))
}

func (s *Solver) PartTwo() {
	filtered := make([]Line, 0)

	for _, l := range s.lines {
		xmin, xmax := lib.FindMinMax([2]int{l.x1, l.x2})
		ymin, ymax := lib.FindMinMax([2]int{l.y1, l.y2})
		if l.x1 == l.x2 || l.y1 == l.y2 || xmax-xmin == ymax-ymin {
			filtered = append(filtered, l)
		}
	}

	fmt.Println("Part one: ", solve(filtered))
}

func generateBoard(lines []Line) map[int]map[int]int {
	board := make(map[int]map[int]int)

	for _, l := range lines {
		xd := 0
		yd := 0
		if l.x1 > l.x2 {
			xd = -1
			l.x2 += -1
		}
		if l.x1 < l.x2 {
			xd = 1
			l.x2 += 1
		}

		if l.y1 > l.y2 {
			yd = -1
			l.y2 += -1
		}
		if l.y1 < l.y2 {
			yd = 1
			l.y2 += 1
		}

		x, y := l.x1, l.y1
		for x != l.x2 || y != l.y2 {
			if _, ok := board[y]; !ok {
				board[y] = make(map[int]int)
			}

			board[y][x]++

			x += xd
			y += yd
		}
	}

	return board
}

func solve(lines []Line) int {
	sea := generateBoard(lines)

	cc := 0
	for _, x := range sea {
		for _, y := range x {
			if y > 1 {
				cc++
			}
		}
	}

	return cc
}
