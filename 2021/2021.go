package main

import (
	"fmt"

	"github.com/jonavdm/aoc/2021/day1"
	"github.com/jonavdm/aoc/2021/day2"
	"github.com/jonavdm/aoc/2021/day3"
)

type Day interface {
	Read()
	PartOne()
	PartTwo()
}

func main() {
	dayCollection := make(map[string]Day)

	dayCollection["day1"] = day1.New("inputs")
	dayCollection["day2"] = day2.New("inputs")
	dayCollection["day3"] = day3.New("inputs")

	for day, obj := range dayCollection {
		fmt.Println(day)
		obj.Read()
		obj.PartOne()
		obj.PartTwo()
		fmt.Println()
	}
}
