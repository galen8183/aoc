package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	input, err := os.ReadFile("test")
	//input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	stones := parse(input)
	solve(stones, 25)
	solve(stones, 75)
}

func solve(stones []int, blinks int) {
	var sum int
	for _, st := range(stones) {
		n := blinkN(st, blinks)
		sum += n
	}
	fmt.Println(sum)
}

var cache = make(map[[2]int]int)

func blinkN(stone, blinks int) int {
	if cached, ok := cache[[2]int{stone, blinks}]; ok {
		return cached
	}

	if blinks == 1 {
		n := len(blink(stone))
		cache[[2]int{stone, blinks}] = n
		return n
	}

	var n int
	for _, st := range(blink(stone)) {
		m := blinkN(st, blinks-1)
		cache[[2]int{st, blinks-1}] = m
		n += m
	}
	return n
}

func blink(st int) []int {
	if st == 0 {
		return []int{1}
	}
	n := strconv.Itoa(st)
	if len(n) % 2 == 0 {
		i, _ := strconv.Atoi(n[:len(n)/2])
		j, _ := strconv.Atoi(n[len(n)/2:])
		return []int{i, j}
	}
	return []int{st * 2024}
}

func parse(input []byte) (stones []int) {
	in := strings.Fields(string(input))
	for _, n := range(in) {
		stone, _ := strconv.Atoi(n)
		stones = append(stones, stone)
	}
	return
}
