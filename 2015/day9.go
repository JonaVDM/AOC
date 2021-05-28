package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

const INPFILE = "day9.input"

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
	locations := createMap(data)
	places := make([]string, 0)

	for place := range locations {
		places = append(places, place)
	}

	var routes [][]string
	generatePermutations(len(locations), places, &routes)

	shortest := math.MaxInt32

	for _, route := range routes {
		distance := 0

		for i := 0; i < len(route)-1; i++ {
			from := route[i]
			to := route[i+1]
			distance += locations[from][to]
		}

		if distance < shortest {
			shortest = distance
		}
	}

	return fmt.Sprint(shortest)
}

func part2(data []string) string {
	locations := createMap(data)
	places := make([]string, 0)

	for place := range locations {
		places = append(places, place)
	}

	var routes [][]string
	generatePermutations(len(locations), places, &routes)

	shortest := math.MinInt32

	for _, route := range routes {
		distance := 0

		for i := 0; i < len(route)-1; i++ {
			from := route[i]
			to := route[i+1]
			distance += locations[from][to]
		}

		if distance > shortest {
			shortest = distance
		}
	}

	return fmt.Sprint(shortest)
}

func createMap(data []string) map[string]map[string]int {
	locations := make(map[string]map[string]int)

	for _, location := range data {
		s := strings.Fields(location)

		distance, _ := strconv.Atoi(s[4])

		if _, e := locations[s[0]]; !e {
			locations[s[0]] = make(map[string]int)
		}

		if _, e := locations[s[2]]; !e {
			locations[s[2]] = make(map[string]int)
		}

		locations[s[0]][s[2]] = distance
		locations[s[2]][s[0]] = distance
	}

	return locations
}

// Heap's Algorithm
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func generatePermutations(n int, strs []string, perms *[][]string) {
	if n == 1 {
		strsCopy := make([]string, len(strs))
		copy(strsCopy, strs)

		*perms = append(*perms, strsCopy)
	} else {
		for i := 0; i < n-1; i++ {
			generatePermutations(n-1, strs, perms)
			if n%2 == 0 {
				swap(strs, i, n-1)
			} else {
				swap(strs, 0, n-1)
			}
		}
		generatePermutations(n-1, strs, perms)
	}
}

func swap(strs []string, i, j int) {
	strs[i], strs[j] = strs[j], strs[i]
}
