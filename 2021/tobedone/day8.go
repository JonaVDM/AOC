package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
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
	counter := 0

	for _, d := range data {
		o := strings.Split(d, " | ")
		for _, i := range strings.Split(o[1], " ") {

			switch len(i) {
			case 2, 4, 3, 7:
				counter++
			}
		}
	}

	return fmt.Sprint(counter)
}

func part2(data []string) string {
	counter := 0
	for _, d := range data {
		inp := strings.Split(d, " | ")
		list := strings.Split(inp[0], " ")
		sort.Slice(list, func(i, j int) bool {
			return len(list[i]) < len(list[j])
		})
		sorted := make([]string, 0)
		for _, l := range list {
			sorted = append(sorted, SortString(l))
		}

		wordMap := make(map[string]string)

		// Things that we know
		wordMap[sorted[0]] = "1"
		wordMap[sorted[1]] = "7"
		wordMap[sorted[2]] = "4"
		wordMap[sorted[9]] = "8"

		posMap := make(map[string]int)
		for _, ins := range sorted {
			for _, l := range strings.Split(ins, "") {
				posMap[l] += 1
			}
		}

		letterMap := make(map[int]string)
		for key, value := range posMap {
			if value == 4 {
				letterMap[5] = key
				continue
			}

			if value == 6 {
				letterMap[2] = key
				continue
			}

			if value == 9 {
				letterMap[6] = key

				if string(sorted[0][0]) == key {
					letterMap[3] = string(sorted[0][1])
				}
				if string(sorted[0][1]) == key {
					letterMap[3] = string(sorted[0][0])
				}
				continue
			}
		}

		for key, value := range posMap {
			if value != 8 || key == letterMap[3] {
				continue
			}

			letterMap[1] = key
		}

		for _, letter := range strings.Split("abcdefg", "") {
			m := true
			for _, value := range letterMap {
				if value == letter {
					m = false
					break
				}
			}
			if m {
				in := false
				for _, l := range strings.Split(sorted[2], "") {
					if letter == l {
						in = true
					}
				}
				if in {
					letterMap[4] = letter
				} else {
					letterMap[7] = letter
				}
			}
		}

		wordMap[SortString(letterMap[1]+letterMap[3]+letterMap[4]+letterMap[5]+letterMap[7])] = "2"
		wordMap[SortString(letterMap[1]+letterMap[3]+letterMap[4]+letterMap[6]+letterMap[7])] = "3"
		wordMap[SortString(letterMap[1]+letterMap[2]+letterMap[4]+letterMap[6]+letterMap[7])] = "5"
		wordMap[SortString(letterMap[1]+letterMap[2]+letterMap[4]+letterMap[5]+letterMap[6]+letterMap[7])] = "6"
		wordMap[SortString(letterMap[1]+letterMap[2]+letterMap[3]+letterMap[4]+letterMap[6]+letterMap[7])] = "9"
		wordMap[SortString(letterMap[1]+letterMap[2]+letterMap[3]+letterMap[5]+letterMap[6]+letterMap[7])] = "0"

		c := ""

		for _, d := range strings.Split(inp[1], " ") {
			c += wordMap[SortString(d)]
		}
		cn, _ := strconv.Atoi(c)
		counter += cn
	}

	return fmt.Sprint(counter)
}

// https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
