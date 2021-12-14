package day3

import (
	"fmt"
	"strconv"
)

func (s *Solver) PartOne() {
	gamma := ""
	epsilon := ""

	for i := range s.data[0] {
		ones := 0
		zeros := 0

		for _, d := range s.data {
			if d[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		if ones > zeros {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println("Part 1: ", g*e)
}

func (s *Solver) PartTwo() {
	mc := make([]string, len(s.data))
	lc := make([]string, len(s.data))

	copy(mc, s.data)
	copy(lc, s.data)

	for i := range s.data[0] {
		mOnes := []string{}
		mZeros := []string{}

		for _, d := range mc {
			if d[i] == '0' {
				mZeros = append(mZeros, d)
			} else {
				mOnes = append(mOnes, d)
			}
		}

		if len(mOnes) >= len(mZeros) {
			mc = mOnes
		} else {
			mc = mZeros
		}
	}

	for i := range s.data[0] {
		if len(lc) == 1 {
			break
		}
		lOnes := []string{}
		lZeros := []string{}

		for _, d := range lc {
			if d[i] == '0' {
				lZeros = append(lZeros, d)
			} else {
				lOnes = append(lOnes, d)
			}
		}

		if len(lOnes) >= len(lZeros) {
			lc = lZeros
		} else {
			lc = lOnes
		}
	}

	o, _ := strconv.ParseInt(mc[0], 2, 64)
	c, _ := strconv.ParseInt(lc[0], 2, 64)

	fmt.Println("Part 2: ", o*c)
}
