package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const INPFILE = "day8.input"

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
	total := 0

	for _, inp := range data {
		copy := inp[1 : len(inp)-1]
		copy = strings.ReplaceAll(copy, `\"`, `Z`)
		copy = strings.ReplaceAll(copy, `\\`, `Z`)

		for strings.Contains(copy, `\x`) {
			i := strings.Index(copy, `\x`)
			copy = copy[:i] + copy[i+3:]
		}

		total += len(inp) - len(copy)
	}

	return fmt.Sprint(total)
}

func part2(data []string) string {
	total := 0

	for _, inp := range data {
		ori := len(inp)

		inp = strings.ReplaceAll(inp, `"`, `ZZ`)
		inp = strings.ReplaceAll(inp, `\`, `ZZ`)
		inp = inp + "ZZ"

		total += len(inp) - ori
	}

	return fmt.Sprint(total)
}
