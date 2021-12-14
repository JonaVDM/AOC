package day7

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func New(folder string) *Solver {
	return &Solver{
		day:    7,
		folder: folder,
	}
}

type Solver struct {
	day      int
	raw      string
	folder   string
	crabs    []int
	smallest int
	biggest  int
}

func (s *Solver) Read() {
	path := fmt.Sprintf("%s/day%d", s.folder, s.day)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s.raw = strings.TrimSpace(string(f))
	s.crabs = make([]int, 0)
	s.smallest = math.MaxInt
	for _, d := range strings.Split(s.raw, ",") {
		n, _ := strconv.Atoi(d)
		s.crabs = append(s.crabs, n)

		if s.smallest > n {
			s.smallest = n
		}

		if s.biggest < n {
			s.biggest = n
		}
	}

}
