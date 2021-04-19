package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type solver struct {
	list   map[string]string
	solved map[string]uint16
}

const INPFILE = "day7.input"

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

func (s *solver) loop(inst string) uint16 {
	if solved, set := s.solved[inst]; set {
		return solved
	}

	parts := strings.Split(s.list[inst], " ")

	if len(parts) == 1 {
		n, err := strconv.Atoi(parts[0])

		if err != nil {
			return s.loop(parts[0])
		}
		return uint16(n)
	}

	if parts[0] == "NOT" {
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return ^s.loop(parts[1])
		}
		return ^uint16(n)
	}

	nt1, err := strconv.Atoi(parts[0])
	var n1 uint16
	if err != nil {
		n1 = s.loop(parts[0])
	} else {
		n1 = uint16(nt1)
	}

	nt2, err := strconv.Atoi(parts[2])
	var n2 uint16
	if err != nil {
		n2 = s.loop(parts[2])
	} else {
		n2 = uint16(nt2)
	}

	if parts[1] == "AND" {
		s.solved[inst] = n1 & n2
	} else if parts[1] == "OR" {
		s.solved[inst] = n1 | n2
	} else if parts[1] == "LSHIFT" {
		s.solved[inst] = n1 << n2
	} else if parts[1] == "RSHIFT" {
		s.solved[inst] = n1 >> n2
	}

	return s.solved[inst]
}

func part1(data []string) string {
	list := make(map[string]string)

	for _, inst := range data {
		p := strings.Split(inst, " -> ")
		list[p[1]] = p[0]
	}

	sl := solver{
		list:   list,
		solved: make(map[string]uint16),
	}
	res := sl.loop("a")
	fmt.Println(res)

	return "Part 1"
}

func part2(data []string) string {
	return "Part 2"
}
