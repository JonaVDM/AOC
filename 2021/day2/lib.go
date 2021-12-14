package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func New(folder string) *Solver {
	return &Solver{
		day:    2,
		folder: folder,
	}
}

type Solver struct {
	day    int
	raw    string
	folder string
	data   []struct {
		Direction string
		amount    int
	}
}

func (s *Solver) Read() {
	path := fmt.Sprintf("%s/day%d", s.folder, s.day)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s.raw = strings.TrimSpace(string(f))
	s.data = make([]struct {
		Direction string
		amount    int
	}, 0)

	for _, d := range strings.Split(s.raw, "\n") {
		ins := strings.Split(d, " ")
		a, _ := strconv.Atoi(ins[1])
		s.data = append(s.data, struct {
			Direction string
			amount    int
		}{
			ins[0],
			a,
		})
	}
}
