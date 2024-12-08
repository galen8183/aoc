package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.Open("test")
	// input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	part1(*scanner)
}

var NumRegex = regexp.MustCompile(`\d+`)

func part1(scanner bufio.Scanner) {
	var t, d []int
	scanner.Scan()
	for _, n := range(NumRegex.FindAll(scanner.Bytes(), -1)) {
		m, _ := strconv.Atoi(string(n))
		t = append(t, m)
	}
	scanner.Scan()
	for _, n := range(NumRegex.FindAll(scanner.Bytes(), -1)) {
		m, _ := strconv.Atoi(string(n))
		d = append(d, m)
	}

	total := 1
	for i := 0; i < len(t); i++ {
		var wins int
		for j := 1; j < t[i]; j++ {
			if (t[i] - j) * j < d[i] + 1 {
				if wins != 0 {
					break
				}
				continue
			}
			wins++
		}
		total *= wins
		wins = 0
	}

	fmt.Println(total)
}
