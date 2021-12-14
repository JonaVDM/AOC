package day1

import (
	"fmt"
	"math"
)

func (s *Solver) PartOne() {
	prev := math.MaxInt
	counter := 0

	for _, d := range s.data {
		if d > prev {
			counter++
		}
		prev = d
	}

	fmt.Println("Part 1: ", counter)
}

func (s *Solver) PartTwo() {
	prev := math.MaxInt
	counter := 0

	for i := 0; i < len(s.data)-2; i++ {
		t := s.data[i] + s.data[i+1] + s.data[i+2]

		if t > prev {
			counter++
		}
		prev = t
	}

	fmt.Println("Part 2: ", counter)
}
