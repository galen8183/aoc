package main

import (
	"fmt"
	"log"
	"os"
)

var chrs [][]byte

func main() {
	//input, err := os.ReadFile("test")
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	var buf []byte
	for _, c := range(input) {
		if c == '\n' {
			chrs = append(chrs, buf)
			buf = []byte{}
			continue
		}
		buf = append(buf, c)
	}
	chrs = append(chrs, buf)
	part1()
	part2()
}

func part1() {
	var total int
	for i, line := range(chrs) {
		for j, c := range(line) {
			if c != 'X' {
				continue
			}

			for n := -1; n < 2; n++ {
				for m := -1; m < 2; m++ {
					if !isChar(i+n, j+m, 'M') {
						continue
					}
					if !isChar(i+n*2, j+m*2, 'A') {
						continue
					}
					if !isChar(i+n*3, j+m*3, 'S') {
						continue
					}
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

func part2() {
	var total int
	for i, line := range(chrs) {
		for j, c := range(line) {
			if c != 'M' {
				continue
			}

			if findInDirection(i, j, 3) {
				if findInDirection(i+2, j, 0) || findInDirection(i, j+2, 2) {
					total++
				}
			}
			if findInDirection(i, j, 1) {
				if findInDirection(i-2, j, 2) || findInDirection(i, j-2, 0) {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

func getInBounds(i, j int) (byte, bool) {
	if i < 0 || i >= len(chrs) {
		return 0, false
	}
	if j < 0 || j >= len(chrs[i]) {
		return 0, false
	}
	return chrs[i][j], true
}

func isChar(i, j int, chr byte) bool {
	c, ok := getInBounds(i, j)
	if !ok {
		return false
	}
	return c == chr
}

func findInDirection(i, j, quad int) bool {
	const word = "MAS"
	var dirs = [][2]int{{-1, 1}, {-1, -1}, {1, -1}, {1, 1}}

	for _, c := range(word) {
		if !isChar(i, j, byte(c)) {
			return false
		}
		i += dirs[quad][0]
		j += dirs[quad][1]
	}

	return true
}
