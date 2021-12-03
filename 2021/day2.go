package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPFILE = "day2.input"

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
	x := 0
	y := 0
	for _, i := range data {
		ins := strings.Split(i, " ")
		a, _ := strconv.Atoi(ins[1])

		if ins[0] == "forward" {
			x += a
		} else if ins[0] == "down" {
			y += a
		} else if ins[0] == "up" {
			y -= a
		} else {
			fmt.Println("Unkown direction", ins[0])
		}
	}

	return fmt.Sprint(x * y)
}

func part2(data []string) string {
	x := 0
	y := 0
	aim := 0

	for _, i := range data {
		ins := strings.Split(i, " ")
		a, _ := strconv.Atoi(ins[1])

		if ins[0] == "forward" {
			x += a
			y += aim * a
		} else if ins[0] == "down" {
			aim += a
		} else if ins[0] == "up" {
			aim -= a
		} else {
			fmt.Println("Unkown direction", ins[0])
		}
	}

	return fmt.Sprint(x * y)
}
