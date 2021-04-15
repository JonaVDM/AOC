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

func move(x *int, y *int, sym string) {
	if sym == "^" {
		// North
		*y += 1
	} else if sym == ">" {
		// East
		*x += 1
	} else if sym == "v" {
		// South
		*y -= 1
	} else if sym == "<" {
		// West
		*x -= 1
	}
}

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	inp := string(f)

	data := strings.Split(inp, "")

	fmt.Printf("Part 1: %s\n", part1(data))
	fmt.Printf("Part 2: %s\n", part2(data))
}

func part1(data []string) string {
	x, y := 0, 0

	vis := make([]location, 0)

	for _, sym := range data {
		if !contains(vis, location{x, y}) {
			vis = append(vis, location{x, y})
		}

		move(&x, &y, sym)
	}

	if !contains(vis, location{x, y}) {
		vis = append(vis, location{x, y})
	}

	return fmt.Sprint(len(vis))
}

func part2(data []string) string {
	xa, ya := 0, 0
	xb, yb := 0, 0

	vis := make([]location, 0)

	for i := 0; i < len(data); i += 2 {
		if !contains(vis, location{xa, ya}) {
			vis = append(vis, location{xa, ya})
		}

		if !contains(vis, location{xb, yb}) {
			vis = append(vis, location{xb, yb})
		}

		move(&xa, &ya, data[i])
		move(&xb, &yb, data[i+1])
	}

	if !contains(vis, location{xa, ya}) {
		vis = append(vis, location{xa, ya})
	}

	if !contains(vis, location{xb, yb}) {
		vis = append(vis, location{xb, yb})
	}

	return fmt.Sprint(len(vis))
}
