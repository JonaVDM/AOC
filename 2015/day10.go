package main

import (
	"fmt"
)

func main() {
	data := "1113122113"
	fmt.Printf("Part 1: %s\n", part1(data))
	fmt.Printf("Part 2: %s\n", part2(data))
}

func part1(data string) string {
	return fmt.Sprint(solver(data, 0, 40))
}

func part2(data string) string {
	return fmt.Sprint(solver(data, 0, 50))
}

func solver(data string, i int, t int) int {
	if i == t {
		return len(data)
	}

	nd := ""

	for len(data) > 0 {
		counter := 1
		number := data[0]

		for i := 1; i < len(data); i++ {
			if data[i] != number {
				break
			}

			counter++
		}

		nd += fmt.Sprintf("%d%c", counter, number)
		data = data[counter:]
	}

	return solver(nd, i+1, t)
}
