package day4

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func New(folder string) *Solver {
	return &Solver{
		day:    4,
		folder: folder,
	}
}

type Solver struct {
	day    int
	raw    string
	folder string

	inputs []string
	boards []Board
}

type Board struct {
	Grid [][]string
	Won  bool
}

func (s *Solver) Read() {
	path := fmt.Sprintf("%s/day%d", s.folder, s.day)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s.raw = strings.TrimSpace(string(f))
	data := strings.Split(s.raw, "\n\n")
	s.inputs = strings.Split(data[0], ",")
	s.boards = make([]Board, 0)

	for _, b := range data[1:] {
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

		s.boards = append(s.boards, Board{
			Grid: grid,
			Won:  false,
		})
	}
}
