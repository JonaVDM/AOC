package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// const INPFILE = "test"

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
	inputs := data[0]
	boards := parseBoards(data[1:])

	for _, inp := range strings.Split(inputs, ",") {
		for _, board := range boards {
			board.Checked[inp] = true

			if board.HasWon() {
				sum := board.GetSum()
				ln, _ := strconv.Atoi(inp)

				return fmt.Sprint(sum * ln)
			}
		}
	}

	return "No board has won"
}

func part2(data []string) string {
	inputs := data[0]
	boards := parseBoards(data[1:])
	bc := len(boards)

	for _, inp := range strings.Split(inputs, ",") {
		for _, board := range boards {
			// For some weird ass reason, updating board.Won to true and than calling
			// that here always returned false... but in a function it works fine
			// WHAT
			if board.HasWon() {
				continue
			}

			board.Checked[inp] = true

			if board.HasWon() {
				if bc > 1 {
					bc--
					continue
				}

				sum := board.GetSum()
				ln, _ := strconv.Atoi(inp)

				return fmt.Sprint(sum * ln)
			}
		}
	}
	return "Jona messed up, fix it and try again"
}

func parseBoards(data []string) []Board {
	boards := []Board{}

	for _, b := range data {
		grid := [][]string{}
		for _, r := range strings.Split(b, "\n") {
			if r == "" {
				continue
			}

			row := []string{}

			for _, c := range strings.Split(r, " ") {
				if c == "" {
					continue
				}
				row = append(row, c)
			}

			grid = append(grid, row)
		}

		boards = append(boards, Board{
			Grid:    grid,
			Checked: make(map[string]bool),
			Won:     false,
		})
	}

	return boards
}

type Board struct {
	Checked map[string]bool
	Grid    [][]string
	Won     bool
}

func (b *Board) HasWon() bool {
	for i, row := range b.Grid {
		cc := 0
		rc := 0
		for _, col := range row {
			if val, ok := b.Checked[col]; !ok || !val {
				break
			}
			rc++
		}

		for _, r := range b.Grid {
			if val, ok := b.Checked[r[i]]; !ok || !val {
				break
			}
			cc++
		}

		if rc == 5 || cc == 5 {
			b.Won = true
			return true
		}
	}

	return false
}

func (b *Board) GetSum() int {
	sum := 0

	for _, row := range b.Grid {
		for _, col := range row {
			if !b.Checked[col] {
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
