package day6

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func New(folder string) *Solver {
	return &Solver{
		day:    6,
		folder: folder,
	}
}

type Solver struct {
	day    int
	raw    string
	folder string
	start  []int
}

func (s *Solver) Read() {
	path := fmt.Sprintf("%s/day%d", s.folder, s.day)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s.raw = strings.TrimSpace(string(f))
	s.start = make([]int, 0)
	for _, d := range strings.Split(s.raw, ",") {
		n, _ := strconv.Atoi(d)
		s.start = append(s.start, n)
	}
}
