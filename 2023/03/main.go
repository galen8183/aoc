package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// input, err := os.Open("test")
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var lines []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// part1(lines)
	part2(lines)
}

func part1(input []string) {
	var sum int
	bounds := [2]int{len(input) - 1, len(input[0]) - 1}
	parts := findNums(input)

OUTER:
	for _, part := range(parts) {
		for i := part.Line - 1; i <= part.Line + 1; i++ {
			for j := part.Start - 1; j <= part.End + 1; j++ {
				c := input[min(max(i, 0), bounds[0])][min(max(j, 0), bounds[1])]
				if (c < '0' || c > '9') && c != '.' {
					// fmt.Printf("%c (%2d,%2d)\n", c, min(max(i, 0), bounds[0]), min(max(j, 0), bounds[1]))
					sum += part.Num
					continue OUTER
				}
			}
		}
	}

	fmt.Println(sum)
}

func part2(input []string) {
	var sum int
	bounds := [2]int{len(input) - 1, len(input[0]) - 1}
	parts := findNums(input)
	var gearParts []PartNum

OUTER:
	for _, part := range(parts) {
		for i := part.Line - 1; i <= part.Line + 1; i++ {
			for j := part.Start - 1; j <= part.End + 1; j++ {
				c := input[min(max(i, 0), bounds[0])][min(max(j, 0), bounds[1])]
				if c == '*' {
					part.Gear = [2]int{i, j}
					gearParts = append(gearParts, part)
					continue OUTER
				}
			}
		}
	}

	for i, part := range(gearParts) {
		for _, part2 := range(gearParts[i:]) {
			if part == part2 {
				continue
			}
			if part.Gear == part2.Gear {
				sum += part.Num * part2.Num
			}
		}
	}

	fmt.Println(sum)
}

type PartNum struct {
	Num   int

	Line  int
	Start int
	End   int
	Gear  [2]int
}

func findNums(input []string) (parts []PartNum) {
	var part PartNum
	var isNum bool

	for i, line := range(input) {
		for j, char := range(line) {
			if char < '0' || char > '9' {
				if isNum {
					part.End = j - 1
					if part.End < part.Start {
						part.End = len(line) - 1
					}

					// awful but necessary debugging :')
					// fmt.Printf("prev %c line %d num %3s [%3d:%3d] at (%3d,%3d)\n", input[part.Line][max(j-1,0)], part.Line, string(line[part.Start:part.End + 1]), part.Start, part.End + 1, i, j)
					num, err := strconv.Atoi(input[part.Line][part.Start:part.End + 1])
					if err != nil {
						log.Fatal(err)
					}

					part.Num = num
					parts = append(parts, part)
					part = PartNum{}
					isNum = false
				}
				continue
			}
			if isNum {
				continue
			}

			part = PartNum{Line: i, Start: j}
			isNum = true
		}
	}

	return
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 1; i < m; i++ {
		result *= n
	}
	return result
}
