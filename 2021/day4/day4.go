package day4

import (
	"fmt"
	"strconv"
)

func (s *Solver) PartOne() {
	checks := make(map[string]bool)
	for _, inp := range s.inputs {
		checks[inp] = true
		for _, board := range s.boards {
			if board.HasWon(checks) {
				sum := board.GetSum(checks)
				ln, _ := strconv.Atoi(inp)

				fmt.Println("Part 1: ", sum*ln)
				return
			}
		}
	}
	fmt.Println("Part 1 has no solution")
}

func (s *Solver) PartTwo() {
	checks := make(map[string]bool)
	bc := len(s.boards)
	for _, inp := range s.inputs {
		checks[inp] = true
		for i, board := range s.boards {
			if s.boards[i].Won {
				continue
			}

			if board.HasWon(checks) {
				s.boards[i].Won = true
				if bc > 1 {
					bc--
					continue
				}

				sum := board.GetSum(checks)
				ln, _ := strconv.Atoi(inp)

				fmt.Println("Part 2: ", sum*ln)
				return
			}
		}
	}
	fmt.Println("Part 2 has no solution")
}

func (b *Board) HasWon(checks map[string]bool) bool {
	if b.Won {
		return true
	}

	for i, row := range b.Grid {
		cc := 0
		rc := 0
		for _, col := range row {
			if val, ok := checks[col]; !ok || !val {
				break
			}
			rc++
		}

		for _, r := range b.Grid {
			if val, ok := checks[r[i]]; !ok || !val {
				break
			}
			cc++
		}

		if rc == 5 || cc == 5 {
			return true
		}
	}

	return false
}

func (b *Board) GetSum(checks map[string]bool) int {
	sum := 0

	for _, row := range b.Grid {
		for _, col := range row {
			if !checks[col] {
				i, _ := strconv.Atoi(col)
				sum += i
			}
		}
	}

	return sum
}

func (b *Board) PlayedOut() bool {
	return b.Won
}
