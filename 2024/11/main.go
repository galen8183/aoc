package main

import (
	"fmt"
	"log"
	"math"
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

	var stones []int
	for _, n := range(strings.Fields(string(input))) {
		m, _ := strconv.Atoi(n)
		stones = append(stones, m)
	}
	solve(stones, 25)
	solve(stones, 75)
}

func solve(st []int, blinks int) {
	fmt.Println(blinkAll(st, blinks))
}

func blinkAll(st []int, blinks int) (sum int) {
	if blinks <= 0 {
		return len(st)
	}

	for _, s := range(st) {
		sum += blinkAll(blink(s), blinks-1)
	}
	return
}

func blink(s int) []int {
	if s == 0 {
		return []int{1}
	}
	n := strconv.Itoa(s)
	if math.Mod(float64(len(n)), 2) == 0 {
		i, _ := strconv.Atoi(n[:len(n)/2])
		j, _ := strconv.Atoi(n[len(n)/2:])
		return []int{i, j}
	}
	return []int{s * 2024}
}
