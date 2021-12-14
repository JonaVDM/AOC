package day7

import (
	"fmt"
	"math"

	"github.com/jonavdm/aoc/lib"
)

func (s *Solver) PartOne() {
	leastFuel := math.MaxInt32

	for i := s.smallest; i <= s.biggest; i++ {
		consumtion := 0
		for _, c := range s.crabs {
			if c > i {
				consumtion += c - i
			} else {
				consumtion += i - c
			}
		}
		if leastFuel > consumtion {
			leastFuel = consumtion
		}
	}
	fmt.Println("Part one: ", leastFuel)
}

func (s *Solver) PartTwo() {
	leastFuel := math.MaxInt32

	for i := s.smallest; i <= s.biggest; i++ {
		consumtion := 0
		for _, c := range s.crabs {
			s, b := lib.MinMaxList([]int{i, c})
			for i := 1; i <= b-s; i++ {
				consumtion += i
				if leastFuel <= consumtion {
					break
				}
			}
			if leastFuel <= consumtion {
				break
			}
		}
		if leastFuel > consumtion {
			leastFuel = consumtion
		}
	}
	fmt.Println("Part two: ", leastFuel)
}
