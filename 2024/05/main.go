package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	input, err := os.Open("test")
	//input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	rules, updates := parse(*scanner)
	part1(rules, updates)
	part2(rules, updates)
}

func parse(scanner bufio.Scanner) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		line := strings.Split(scanner.Text(), "|")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])

		rules[x] = append(rules[x], y)
	}

	var updates [][]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		var buf []int
		for _, p := range(line) {
			n, _ := strconv.Atoi(p)
			buf = append(buf, n)
		}
		updates = append(updates, buf)
	}

	return rules, updates
}

func part1(rules map[int][]int, updates [][]int) {
	var sum int
	for _, update := range(updates) {
		if ok := checkUpdate(rules, update); ok {
			sum += update[len(update) / 2]
		}
	}

	fmt.Println(sum)
}

func part2(rules map[int][]int, updates [][]int) {
	var sum int
	for _, update := range(updates) {
		if ok := checkUpdate(rules, update); ok {
			continue
		}

		var sorted []int
		for len(sorted) < len(update) {
			for _, n := range(update) {
				if _, ok := indexOf(n, sorted); ok {
					continue
				}
				sorted = visit(n, rules, sorted, update)
			}
		}

		sum += sorted[len(sorted) / 2]
	}

	fmt.Println(sum)
}

func indexOf(e int, s []int) (int, bool) {
	for i, el := range(s) {
		if e == el {
			return i, true
		}
	}
	return 0, false
}

func checkUpdate(rules map[int][]int, update []int) bool {
	for r, rule := range(rules) {
		x, ok := indexOf(r, update)
		if !ok {
			continue
		}

		left := len(update)
		for _, n := range(rule) {
			m, ok := indexOf(n, update)
			if ok && m < left {
				left = m
			}
		}

		if left < x {
			return false
		}
	}

	return true
}

func visit(n int, rules map[int][]int, sorted, update []int) []int {
	if _, ok := indexOf(n, sorted); ok {
		return sorted
	}

	for _, m := range(rules[n]) {
		if _, ok := indexOf(m, update); !ok {
			continue
		}
		sorted = visit(m, rules, sorted, update)
	}

	return append([]int{n}, sorted...)
}
