package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const INPFILE = "day14.input"

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	inp := string(f)
	inp = strings.TrimSpace(inp)

	data := strings.Split(inp, "\n\n")

	fmt.Printf("Part 1: %d\n", part1(data[0], strings.Split(data[1], "\n")))
	fmt.Printf("Part 2: %d\n", part2(data[0], strings.Split(data[1], "\n")))
}

func part1(start string, data []string) int64 {
	instructions := parse(data)

	return run(start, instructions, 10)
}

func part2(start string, data []string) int64 {
	instructions := parse(data)

	return run(start, instructions, 40)
}

func parse(data []string) map[string]string {
	rec := make(map[string]string)

	for _, d := range data {
		set := strings.Split(d, " -> ")
		rec[set[0]] = set[1]
	}

	return rec
}

func run(start string, instructions map[string]string, amount int) int64 {
	letters := make(map[string]int64)
	pairs := make(map[string]int64)

	letters[string(start[0])] = 1
	for i := 0; i < len(start)-1; i++ {
		letters[string(start[i+1])] += 1
		pairs[start[i:i+2]] += 1
	}

	for i := 0; i < amount; i++ {
		newPair := make(map[string]int64)
		for pair, amo := range pairs {
			l := instructions[pair]
			letters[l] += amo
			newPair[pair[:1]+l] += amo
			newPair[l+pair[1:]] += amo
		}
		pairs = newPair
	}

	var big, small int64 = 0, math.MaxInt64
	for _, a := range letters {
		if a > big {
			big = a
		}

		if a < small {
			small = a
		}
	}

	return big - small
}
