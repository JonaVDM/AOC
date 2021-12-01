package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPFILE = "day1.input"

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
	prev := 1000000000000000000
	counter := 0

	for _, d := range data {
		n, _ := strconv.Atoi(d)

		if n > prev {
			counter++
		}
		prev = n
	}

	return fmt.Sprint(counter)
}

func part2(data []string) string {
	prev := 1000000000000000000
	counter := 0

	for i := 0; i < len(data)-2; i++ {
		n1, _ := strconv.Atoi(data[i])
		n2, _ := strconv.Atoi(data[i+1])
		n3, _ := strconv.Atoi(data[i+2])

		s := n1 + n2 + n3

		if s > prev {
			counter++
		}

		prev = s
	}

	return fmt.Sprint(counter)
}
