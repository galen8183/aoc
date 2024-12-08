package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("test")
	//input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	//part1(*scanner)
	part2(*scanner)
}

func part1(scanner bufio.Scanner) {
	var seeds []int
	if scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		for i, seed := range(line) {
			if i == 0 {
				continue
			}
			n, err := strconv.Atoi(seed)
			if err != nil {
				log.Fatal(err)
			}
			seeds = append(seeds, n)
		}
	}

	if scanner.Scan() && scanner.Scan() {
		maps := parse(scanner)
		for _, m := range(maps) {
			for i, n := range(seeds) {
				for _, r := range(m) {
					if n < r[1] || n - r[2] > r[1] {
						continue
					}

					seeds[i] = n - (r[1] - r[0])
					break
				}
			}
		}

		var low int
		for _, n := range(seeds) {
			if n == 0 {
				continue
			}
			if low == 0 {
				low = n
			}
			low = min(low, n)
		}
		fmt.Println(low)
	}
}

func part2(scanner bufio.Scanner) {
	var seeds [][2]int
	if scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")[1:]
		for i := 0; i < (len(line) / 2); i++ {
			start, err := strconv.Atoi(line[i * 2])
			if err != nil {
				log.Fatal(err)
			}
			length, err := strconv.Atoi(line[i * 2 + 1])
			if err != nil {
				log.Fatal(err)
			}

			seeds = append(seeds, [2]int{start, start + length})
		}
	}

	if scanner.Scan() && scanner.Scan() {
		maps := parse(scanner)
		for _, m := range(maps) {
			for i, n := range(seeds) {
				for _, r := range(m) {
					if  n[1] <= r[1] || n[0] - r[2] >= r[1] {
						continue
					}

					seeds = append(seeds, [2]int{
						max(max(n[0], r[1]) - (r[1] - r[0]), seeds[i][1]),
						min(min(n[1], r[1] + r[2]) - (r[1] - r[0]), seeds[i][0]),
					})
					break
				}
			}
		}

		var low int
		for _, n := range(seeds) {
			fmt.Println(n)
			if n[0] == 0 {
				continue
			}
			if low == 0 {
				low = n[0]
			}
			low = min(low, n[0])
		}
		fmt.Println(len(seeds), low)
	}
}

type Range [3]int
type Map []Range

func parse(scanner bufio.Scanner) (maps []Map) {
	var m Map
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			maps = append(maps, m)
			m = Map{}
			if scanner.Scan() {
				_ = scanner.Scan()
			}
		}

		var elem Range
		vals := strings.Split(scanner.Text(), " ")
		for i, seed := range(vals) {
			n, err := strconv.Atoi(seed)
			if err != nil {
				log.Fatal(err)
			}
			elem[i] = n
		}
		m = append(m, elem)
	}

	maps = append(maps, m)
	return
}
