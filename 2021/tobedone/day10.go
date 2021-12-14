package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const INPFILE = "day10.input"

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
	counter := 0
	for _, line := range data {
		ill := checkLine(line)

		if ill == ')' {
			counter += 3
		}

		if ill == ']' {
			counter += 57
		}

		if ill == '}' {
			counter += 1197
		}

		if ill == '>' {
			counter += 25137
		}
	}

	return fmt.Sprint(counter)
}

func part2(data []string) string {
	counters := make([]int, 0)

	for _, line := range data {
		ill := checkLine(line)

		if ill != ' ' {
			continue
		}

		comp := completeLine(line)

		counters = append(counters, countLine(comp))
	}

	l := len(counters) / 2
	sort.Ints(counters)

	return fmt.Sprint(counters[l])
}

func checkLine(line string) rune {
	opened := make([]rune, 0)

	for _, cha := range line {
		if cha == '(' || cha == '[' || cha == '{' || cha == '<' {
			opened = append(opened, cha)
			continue
		}
		l := len(opened) - 1

		if cha == ')' && opened[l] == '(' {
			opened = opened[:l]
			continue
		}

		if cha == ']' && opened[l] == '[' {
			opened = opened[:l]
			continue
		}

		if cha == '}' && opened[l] == '{' {
			opened = opened[:l]
			continue
		}

		if cha == '>' && opened[l] == '<' {
			opened = opened[:l]
			continue
		}

		return cha
	}

	return ' '
}

func completeLine(line string) []rune {
	opened := make([]rune, 0)

	for _, cha := range line {
		if cha == '(' || cha == '[' || cha == '{' || cha == '<' {
			opened = append(opened, cha)
			continue
		}
		l := len(opened) - 1

		if cha == ')' && opened[l] == '(' {
			opened = opened[:l]
		}

		if cha == ']' && opened[l] == '[' {
			opened = opened[:l]
		}

		if cha == '}' && opened[l] == '{' {
			opened = opened[:l]
		}

		if cha == '>' && opened[l] == '<' {
			opened = opened[:l]
		}
	}

	return reverseArray(opened)
}

func countLine(line []rune) int {
	c := 0

	for _, cha := range line {
		c *= 5

		if cha == '(' {
			c += 1
			continue
		}

		if cha == '[' {
			c += 2
			continue
		}

		if cha == '{' {
			c += 3
			continue
		}

		if cha == '<' {
			c += 4
			continue
		}
	}

	return c
}

// https://www.tutorialspoint.com/write-a-golang-program-to-reverse-an-array
func reverseArray(arr []rune) []rune {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
