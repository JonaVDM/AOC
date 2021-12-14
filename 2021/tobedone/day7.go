package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

const INPFILE = "day7.input"

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
	crabs := make([]int, 0)

	for _, d := range data {
		n, _ := strconv.Atoi(d)
		crabs = append(crabs, n)
	}

	leastFuel := math.MaxInt32
	smallest, biggest := getMinMax(crabs)

	for i := smallest; i <= biggest; i++ {
		consumtion := 0
		for _, c := range crabs {
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

	return fmt.Sprint(leastFuel)
}

func part2(data []string) string {
	crabs := make([]int, 0)

	for _, d := range data {
		n, _ := strconv.Atoi(d)
		crabs = append(crabs, n)
	}

	leastFuel := math.MaxInt32
	smallest, biggest := getMinMax(crabs)

	for i := smallest; i <= biggest; i++ {
		consumtion := 0
		for _, c := range crabs {
			s, b := getMinMax([]int{i, c})
			for i := 1; i <= b-s; i++ {
				consumtion += i
			}
		}
		if leastFuel > consumtion {
			leastFuel = consumtion
		}
	}

	return fmt.Sprint(leastFuel)
}

func getMinMax(nums []int) (int, int) {
	smallest := nums[0]
	biggest := nums[0]
	for _, n := range nums {
		if smallest > n {
			smallest = n
		}

		if biggest < n {
			biggest = n
		}
	}

	return smallest, biggest
}
