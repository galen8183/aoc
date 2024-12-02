package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"strconv"
)

func main() {
	//input, err := os.Open("test")
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	reps := parse(*scanner)
	//part1(reps)
	part2(reps)
}

func parse(scanner bufio.Scanner) (reps [][]int) {
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var rep []int
		for _, n := range(line) {
			nn, _ := strconv.Atoi(n)
			rep = append(rep, nn)
		}

		reps = append(reps, rep)
	}
	return
}

func part1(reps [][]int) {
	var safe int
OUTER:
	for _, rep := range(reps) {
		var last int
		var inc bool
		for i, lvl := range(rep) {
			if i == 0 {
				last = lvl
				continue
			}
			diff := math.Abs(float64(last - lvl))
			if diff < 1 || diff > 3 {
				continue OUTER
			}
			if i == 1 {
				inc = lvl > last
			} else if inc && lvl < last {
				continue OUTER
			} else if !inc && lvl > last {
				continue OUTER
			}
			last = lvl
		}

		safe++
	}

	fmt.Println(safe)
}

func part2(reps [][]int) {
	var safe int
OUTER:
	for _, rep := range(reps) {
		log.Println("checking", rep)
		var last, unsafe int
		var inc bool
		for i, lvl := range(rep) {
			if i == 0 {
				last = lvl
				continue
			}
			diff := math.Abs(float64(last - lvl))
			if diff < 1 || diff > 3 {
				log.Println("bad diff")
				unsafe++
			}
			if i == 1 {
				inc = lvl > last
			} else if inc && lvl < last {
				log.Println("should inc, got", last, lvl)
				unsafe++
			} else if !inc && lvl > last {
				log.Println("should dec, got", last, lvl)
				unsafe++
			}
			last = lvl
		}

		if unsafe > 1 {
			log.Println("unsafe, skipping")
			continue OUTER
		}

		log.Println("safe report")
		safe++
		unsafe = 0
	}

	fmt.Println(safe)
}
