package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const INPFILE = "day3.input"

type location struct {
	x int
	y int
}

func contains(s []location, e location) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	data := string(f)

	fmt.Printf("Part 1: %s\n", part1(data))
	fmt.Printf("Part 2: %s\n", part2(data))
}

func part1(data string) string {
	x, y := 0, 0

	vis := make([]location, 0)

	for _, sym := range strings.Split(data, "") {
		if !contains(vis, location{x, y}) {
			vis = append(vis, location{x, y})
		}

		if sym == "^" {
			// North
			y += 1
		} else if sym == ">" {
			// East
			x += 1
		} else if sym == "v" {
			// South
			y -= 1
		} else if sym == "<" {
			// West
			x -= 1
		}
	}

	return fmt.Sprint(len(vis))
}
