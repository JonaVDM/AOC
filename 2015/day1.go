package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const INPFILE = "day1.input"

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
	list := strings.Split(data, "")
	floor := 0
	for i := range list {
		if list[i] == "(" {
			floor += 1
		} else {
			floor -= 1
		}
	}
	fmt.Println(floor)
	return fmt.Sprintf("%d", floor)
}

func part2(data string) string {
	return "Part 2"
}
