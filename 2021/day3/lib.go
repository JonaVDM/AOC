package day3

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func New(folder string) *Solver {
	return &Solver{
		day:    3,
		folder: folder,
	}
}

type Solver struct {
	day    int
	raw    string
	folder string
	data   []string
}

func (s *Solver) Read() {
	path := fmt.Sprintf("%s/day%d", s.folder, s.day)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s.raw = strings.TrimSpace(string(f))
	s.data = strings.Split(s.raw, "\n")
}
