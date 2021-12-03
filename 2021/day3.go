package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPFILE = "day3.input"

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
	gamma := ""
	epsilon := ""

	for i := range data[0] {
		ones := 0
		zeros := 0

		for _, d := range data {
			if d[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		if ones > zeros {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return fmt.Sprint(g * e)
}

func part2(data []string) string {
	mc := make([]string, len(data))
	lc := make([]string, len(data))

	copy(mc, data)
	copy(lc, data)

	for i := range data[0] {
		mOnes := []string{}
		mZeros := []string{}

		for _, d := range mc {
			if d[i] == '0' {
				mZeros = append(mZeros, d)
			} else {
				mOnes = append(mOnes, d)
			}
		}

		if len(mOnes) >= len(mZeros) {
			mc = mOnes
		} else {
			mc = mZeros
		}
	}

	for i := range data[0] {
		if len(lc) == 1 {
			break
		}
		lOnes := []string{}
		lZeros := []string{}

		for _, d := range lc {
			if d[i] == '0' {
				lZeros = append(lZeros, d)
			} else {
				lOnes = append(lOnes, d)
			}
		}

		if len(lOnes) >= len(lZeros) {
			lc = lZeros
		} else {
			lc = lOnes
		}
	}

	o, _ := strconv.ParseInt(mc[0], 2, 64)
	c, _ := strconv.ParseInt(lc[0], 2, 64)

	return fmt.Sprint(o * c)
}
