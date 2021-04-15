package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const INPFILE = "day5.input"

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
	vowels := "aeiou"
	nice := 0
	forbidden := []string{"ab", "cd", "pq", "xy"}

	for _, row := range data {
		vCount := 0
		dCount := 0
		nope := false

		for i, letter := range row {
			if i < len(row)-1 {
				if row[i] == row[i+1] {
					dCount++
				}

				for _, fb := range forbidden {
					if row[i:i+2] == fb {
						fmt.Println("Hello")
						nope = true
					}
				}
			}

			for _, vowel := range vowels {
				if letter == vowel {
					vCount++
				}
			}
		}

		if nope {
			continue
		}

		if vCount < 3 {
			continue
		}

		if dCount < 1 {
			continue
		}

		nice++
	}

	return fmt.Sprint(nice)
}

func part2(data []string) string {
	return "Part 2"
}
