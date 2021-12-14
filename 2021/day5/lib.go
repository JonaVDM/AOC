package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func New(folder string) *Solver {
	return &Solver{
		day:    5,
		folder: folder,
	}
}

type Solver struct {
	day    int
	raw    string
	folder string
	lines  []Line
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (s *Solver) Read() {
	path := fmt.Sprintf("%s/day%d", s.folder, s.day)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s.raw = strings.TrimSpace(string(f))
	s.lines = make([]Line, 0)
	t := regexp.MustCompile(`(,|->)`)
	for _, l := range strings.Split(s.raw, "\n") {
		values := t.Split(l, -1)

		n := make([]int, 0)
		for _, v := range values {
			i, _ := strconv.Atoi(strings.TrimSpace(v))
			n = append(n, i)
		}

		s.lines = append(s.lines, Line{
			n[0],
			n[1],
			n[2],
			n[3],
		})
	}
}
