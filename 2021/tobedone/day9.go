package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

const INPFILE = "day9.input"

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
	caves := parse(data)

	lower := 0
	for y := range caves {
		for x, curr := range caves[y] {
			if y > 0 && caves[y-1][x] <= curr {
				continue
			}

			if x > 0 && caves[y][x-1] <= curr {
				continue
			}

			if y < len(caves)-1 && caves[y+1][x] <= curr {
				continue
			}

			if x < len(caves[y])-1 && caves[y][x+1] <= curr {
				continue
			}

			lower += curr + 1
		}
	}

	return fmt.Sprint(lower)
}

func part2(data []string) string {
	caves := parse(data)

	s := System{
		Caves: caves,
	}

	basins := s.FindBasins()
	sort.Ints(basins)
	biggest := basins[len(basins)-3:]

	return fmt.Sprint(biggest[0] * biggest[1] * biggest[2])
}

func parse(data []string) [][]int {
	caves := make([][]int, 0)

	for _, row := range data {
		r := make([]int, 0)
		for _, col := range strings.Split(row, "") {
			n, _ := strconv.Atoi(col)
			r = append(r, n)
		}
		caves = append(caves, r)
	}

	return caves
}

type System struct {
	Caves [][]int
	been  map[int]map[int]bool
}

func (s *System) FindBasins() []int {
	basins := make([]int, 0)
	s.been = make(map[int]map[int]bool)

	for y := range s.Caves {
		for x := range s.Caves[y] {
			if been, ok := s.been[y][x]; s.Caves[y][x] == 9 || (ok && been) {
				continue
			}

			out := s.discoverBasin(x, y)
			basins = append(basins, out)
		}
	}

	return basins
}

func (s *System) discoverBasin(x, y int) int {
	if _, ok := s.been[y]; !ok {
		s.been[y] = make(map[int]bool)
	}

	if been, ok := s.been[y][x]; ok && been {
		return 0
	}

	s.been[y][x] = true

	if s.Caves[y][x] == 9 {
		return 0
	}

	yp := 0
	xp := 0
	ym := 0
	xm := 0

	if y > 0 {
		ym = s.discoverBasin(x, y-1)
	}

	if x > 0 {
		xm = s.discoverBasin(x-1, y)
	}

	if y < len(s.Caves)-1 {
		yp = s.discoverBasin(x, y+1)
	}

	if x < len(s.Caves[y])-1 {
		xp = s.discoverBasin(x+1, y)
	}

	return 1 + yp + xp + ym + xm
}
