package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const INPFILE = "day12.input"

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
	tunnels := parse(data)

	caves := search(tunnels.Tunnels, "start")

	return fmt.Sprint(len(caves))
}

func part2(data []string) string {
	tunnels := parse(data)

	caves := searchBetter(tunnels.Tunnels, "start", false)

	checked := make(map[string]int)

	for _, c := range caves {
		checked[c] += 1
	}

	return fmt.Sprint(len(checked))
}

func parse(data []string) Cave {
	tunnels := make(map[string][]string)

	for _, line := range data {
		tunnel := strings.Split(line, "-")

		if _, ok := tunnels[tunnel[0]]; !ok {
			tunnels[tunnel[0]] = make([]string, 0)
		}

		if _, ok := tunnels[tunnel[1]]; !ok {
			tunnels[tunnel[1]] = make([]string, 0)
		}

		tunnels[tunnel[0]] = append(tunnels[tunnel[0]], tunnel[1])
		tunnels[tunnel[1]] = append(tunnels[tunnel[1]], tunnel[0])
	}

	return Cave{
		Tunnels: tunnels,
	}
}

type Cave struct {
	Tunnels map[string][]string
}

func search(cave map[string][]string, current string) []string {
	pos := make([]string, 0)
	if current == "end" {
		return []string{current}
	}

	newCave := copyCave(cave)
	if current == strings.ToLower(current) {
		delete(newCave, current)
	}

	for _, tunnel := range cave[current] {
		nextTunnels := search(newCave, tunnel)
		for _, n := range nextTunnels {
			pos = append(pos, fmt.Sprintf("%s,%s", current, n))
		}
	}

	return pos
}

func copyCave(cave map[string][]string) map[string][]string {
	newCave := make(map[string][]string)
	for k := range cave {
		newCave[k] = make([]string, len(cave[k]))
		copy(newCave[k], cave[k])
	}

	return newCave
}

func searchBetter(cave map[string][]string, current string, double bool) []string {
	pos := make([]string, 0)
	if current == "end" {
		return []string{current}
	}

	isLower := false
	newCave := copyCave(cave)
	if current == strings.ToLower(current) {
		delete(newCave, current)
		isLower = true
	}

	for _, tunnel := range cave[current] {
		nextTunnels := searchBetter(newCave, tunnel, double)
		for _, n := range nextTunnels {
			pos = append(pos, fmt.Sprintf("%s,%s", current, n))
		}

		if double || !isLower || current == "start" {
			continue
		}
		// break

		moreTunnels := searchBetter(cave, tunnel, true)
		for _, n := range moreTunnels {
			if n != "end" {
				pos = append(pos, fmt.Sprintf("%s,%s", current, n))
			}
		}
	}

	return pos
}
