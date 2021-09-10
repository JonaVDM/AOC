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
	correct := 0

	for _, line := range data {
		d := strings.Split(line, " ")

		minmax := strings.Split(d[0], "-")

		// Discarding errors, yolo
		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])
		letter := d[1][:1]
		password := d[2]

		c := strings.Count(password, letter)
		if c < min || c > max {
			continue
		}

		correct++
	}

	return fmt.Sprint(correct)
}

func part2(data []string) string {
	correct := 0

	for _, line := range data {
		d := strings.Split(line, " ")

		minmax := strings.Split(d[0], "-")

		// Discarding errors, yolo
		n1, _ := strconv.Atoi(minmax[0])
		n2, _ := strconv.Atoi(minmax[1])
		letter := []byte(d[1])[0]
		password := d[2]

		if (password[n1-1] == letter || password[n2-1] == letter) && password[n1-1] != password[n2-1] {
			correct++
		}
	}

	return fmt.Sprint(correct)
}
