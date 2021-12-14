package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const INPFILE = "day5.input"

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
	lines := parse(data)
	filtered := make([]Line, 0)

	for _, l := range lines {
		if l.x1 == l.x2 || l.y1 == l.y2 {
			filtered = append(filtered, l)
		}
	}

	board := Board{
		Lines: filtered,
	}

	sea := board.Generate()

	cc := 0
	for _, x := range sea {
		for _, y := range x {
			if y > 1 {
				cc++
			}
		}
	}

	return fmt.Sprint(cc)
}

func part2(data []string) string {
	lines := parse(data)
	filtered := make([]Line, 0)

	for _, l := range lines {
		xmin, xmax := findMinAndMax([2]int{l.x1, l.x2})
		ymin, ymax := findMinAndMax([2]int{l.y1, l.y2})
		if l.x1 == l.x2 || l.y1 == l.y2 || xmax-xmin == ymax-ymin {
			filtered = append(filtered, l)
		}
	}

	board := Board{
		Lines: filtered,
	}

	sea := board.Generate()

	cc := 0
	for _, x := range sea {
		for _, y := range x {
			if y > 1 {
				cc++
			}
		}
	}

	return fmt.Sprint(cc)
}

func parse(data []string) []Line {
	lines := make([]Line, 0)
	t := regexp.MustCompile(`(,|->)`)

	for _, l := range data {
		values := t.Split(l, -1)

		n := make([]int, 0)
		for _, v := range values {
			i, _ := strconv.Atoi(strings.TrimSpace(v))
			n = append(n, i)
		}

		lines = append(lines, Line{
			n[0],
			n[1],
			n[2],
			n[3],
		})
	}

	return lines
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Board struct {
	Lines []Line
}

func (b *Board) Generate() map[int]map[int]int {
	board := make(map[int]map[int]int)

	for _, l := range b.Lines {
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

// copied from https://learningprogramming.net/golang/golang-golang/find-max-and-min-of-array-in-golang/
func findMinAndMax(a [2]int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
