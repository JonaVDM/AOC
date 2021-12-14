package day6

import (
	"fmt"
)

func (s *Solver) PartOne() {
	fmt.Println("Part one: ", s.simulator(80))
}

func (s *Solver) PartTwo() {
	fmt.Println("Part two: ", s.simulator(256))
}

func (s *Solver) simulator(days int) int {
	// After looking quite a bit on reddit I found that quite a few
	// people implemented something that looks like this.
	sea := make([]int, 9)

	for _, d := range s.start {
		sea[d]++
	}

	for i := 0; i < days; i++ {
		s0 := sea[0]
		for i := range sea {
			if i == 8 {
				break
			}

			sea[i] = sea[i+1]
		}

		sea[6] += s0
		sea[8] = s0
	}
	total := 0
	for _, fishes := range sea {
		total += fishes
	}
	return total
}
