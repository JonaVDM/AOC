package main

import (
	"fmt"

	"github.com/jonavdm/aoc/2021/day1"
	"github.com/jonavdm/aoc/2021/day2"
	"github.com/jonavdm/aoc/2021/day3"
	"github.com/jonavdm/aoc/2021/day4"
	"github.com/jonavdm/aoc/2021/day5"
)

type Day interface {
	Read()
	PartOne()
	PartTwo()
}

func main() {
	inputFolder := "inputs"
	dayCollection := make([]Day, 0)

	dayCollection = append(dayCollection, day1.New(inputFolder))
	dayCollection = append(dayCollection, day2.New(inputFolder))
	dayCollection = append(dayCollection, day3.New(inputFolder))
	dayCollection = append(dayCollection, day4.New(inputFolder))
	dayCollection = append(dayCollection, day5.New(inputFolder))

	for day, obj := range dayCollection {
		fmt.Printf("-- Day %d --\n", day+1)
		obj.Read()
		obj.PartOne()
		obj.PartTwo()
		fmt.Println()
	}
}
