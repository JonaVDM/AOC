package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// const INPFILE = "test"

const INPFILE = "day6.input"

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	inp := string(f)

	inp = strings.TrimSpace(inp)
	data := strings.Split(inp, ",")

	fmt.Printf("Part 1: %s\n", part1(data))
	fmt.Printf("Part 2: %s\n", part2(data))
}

func part1(data []string) string {
	sea := make([]*Fish, 0)

	for _, d := range data {
		n, _ := strconv.Atoi(d)
		sea = append(sea, &Fish{
			Age: n,
		})
	}

	for i := 1; i <= 80; i++ {
		newFish := make([]*Fish, 0)

		for f := range sea {
			if sea[f].Timer() {
				newFish = append(newFish, &Fish{
					Age: 8,
				})
			}
		}

		sea = append(sea, newFish...)
	}

	return fmt.Sprint(len(sea))
}

func part2(data []string) string {
	// After looking quite a bit on reddit I found that quite a few
	// people implemented something that looks like this.
	sea := make([]int, 9)

	for _, d := range data {
		n, _ := strconv.Atoi(d)
		sea[n]++
	}

	for i := 0; i < 256; i++ {
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
	return fmt.Sprint(total)
}

type Fish struct {
	Age int
}

func (f *Fish) Timer() bool {
	f.Age -= 1

	if f.Age < 0 {
		f.Age = 6
		return true
	}

	return false
}
