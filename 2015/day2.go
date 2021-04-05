package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

const INPFILE = "day2.input"

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	data := string(f)
	data = data[:len(data)-1]

	fmt.Printf("Part 1: %s\n", part1(data))
	fmt.Printf("Part 2: %s\n", part2(data))
}

func part1(data string) string {
	total := 0
	for _, i := range strings.Split(data, "\n") {
		box := getSides(i)
		total += box[0] + box[0]*2 + box[1]*2 + box[2]*2
	}

	return fmt.Sprint(total)
}

func part2(data string) string {
	return "Part 2"
}

func getSides(box string) []int {
	dim := []int{}
	for _, i := range strings.Split(box, "x") {
		i, _ := strconv.Atoi(i)
		dim = append(dim, i)
	}

	sides := make([]int, 3)
	sides[0] = dim[0] * dim[1]
	sides[1] = dim[1] * dim[2]
	sides[2] = dim[0] * dim[2]

	sort.Ints(sides)
	return sides
}
