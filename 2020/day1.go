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

	rawData := strings.Split(inp, "\n")
	rawData = rawData[:len(rawData)-1]

	data := []int{}
	for _, i := range rawData {
		// discarding the error, yolo
		n, _ := strconv.Atoi(i)
		data = append(data, n)
	}

	fmt.Printf("Part 1: %s\n", part1(data))
	fmt.Printf("Part 2: %s\n", part2(data))
}

func part1(data []int) string {
	for _, n1 := range data {
		for _, n2 := range data {
			if n1+n2 != 2020 {
				continue
			}

			return fmt.Sprint(n1 * n2)
		}
	}

	return "Solution not found"
}

func part2(data []int) string {
	for _, n1 := range data {
		for _, n2 := range data {
			for _, n3 := range data {
				if n1+n2+n3 != 2020 {
					continue
				}

				return fmt.Sprint(n1 * n2 * n3)
			}
		}
	}

	return "Solution not found"
}
