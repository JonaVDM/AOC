package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPFILE = "day13.input"

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	inp := string(f)

	data := strings.Split(inp, "\n")
	data = data[:len(data)-1]

	fmt.Printf("Part 1: %s\n", part1(data))
	fmt.Printf("Part 2: %s\n", part2(data))
}

func part1(data []string) string {
	busses := make([]int, 0)
	for _, b := range strings.Split(data[1], ",") {
		if b == "x" {
			continue
		}
		// There can't be anything wrong with the input right?
		n, _ := strconv.Atoi(b)

		busses = append(busses, n)
	}

	// Again just asume the user didn't mess up anything here okay?
	start, _ := strconv.Atoi(data[0])
	for time := start; true; time++ {
		for _, bus := range busses {
			if time%bus != 0 {
				continue
			}

			return fmt.Sprint((time - start) * bus)
		}
	}

	return "we can't get here bus sure go"
}

func part2(data []string) string {
	return "Part 2"
}
