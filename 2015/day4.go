package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

const INPFILE = "day4.input"

func main() {
	f, err := ioutil.ReadFile(INPFILE)
	if err != nil {
		log.Fatal(err)
	}

	key := string(f)
	key = key[:len(key)-1]

	fmt.Printf("Part 1: %s\n", part1(key))
	fmt.Printf("Part 2: %s\n", part2(key))
}

func part1(key string) string {
	i := 0
	for {
		i++
		data := []byte(fmt.Sprintf("%s%d", key, i))
		hash := md5.Sum(data)
		str := hex.EncodeToString(hash[:])
		if str[:5] == "00000" {
			return fmt.Sprint(i)
		}
	}
}

func part2(key string) string {
	i := 0
	for {
		i++
		data := []byte(fmt.Sprintf("%s%d", key, i))
		hash := md5.Sum(data)
		str := hex.EncodeToString(hash[:])
		if str[:6] == "000000" {
			return fmt.Sprint(i)
		}
	}
}
