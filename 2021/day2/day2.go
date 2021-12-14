package day2

import (
	"fmt"
)

func (s *Solver) PartOne() {
	x := 0
	y := 0
	for _, ins := range s.data {
		if ins.Direction == "forward" {
			x += ins.amount
		} else if ins.Direction == "down" {
			y += ins.amount
		} else if ins.Direction == "up" {
			y -= ins.amount
		}
	}

	fmt.Println("Part 1: ", x*y)
}

func (s *Solver) PartTwo() {
	x := 0
	y := 0
	aim := 0
	for _, ins := range s.data {
		if ins.Direction == "forward" {
			x += ins.amount
			y += aim * ins.amount
		} else if ins.Direction == "down" {
			aim += ins.amount
		} else if ins.Direction == "up" {
			aim -= ins.amount
		}
	}

	fmt.Println("Part 1: ", x*y)
}
