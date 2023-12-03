package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	"strings"
)

func main() {
	// input, err := os.Open("test")
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	// part1(*scanner)
	part2(*scanner)
}

func part1(scanner bufio.Scanner) {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		for _, c := range(line) {
			if c < '0' || c > '9' {
				continue
			}

			// fmt.Print(int(c) - '0', " ")
			sum += (int(c) - '0') * 10
			break
		}

		for i := 1; i <= len(line); i++ {
			c := line[len(line) - i]
			if c < '0' || c > '9' {
				continue
			}

			// fmt.Print(int(c) - '0', " ")
			sum += int(c) - '0'
			break
		}

		// fmt.Println(line)
	}

	fmt.Println(sum)
}

// var nums = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var nums = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func part2(scanner bufio.Scanner) {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		/*
		min := 0
		minIdx := len(line)
		max := 0
		maxIdx := 0
		for i, num := range(nums) {
			x := strings.Index(line, num)
			if x == -1 {
				continue
			}

			if x < minIdx {
				min = i
				minIdx = x
			}

			y := strings.LastIndex(line, num)
			if y >= maxIdx {
				max = i
				maxIdx = y
			}
		}

		fmt.Printf("min %2d, max %2d, line %s\n", int(math.Mod(float64(min), 9.0)) + 1, int(math.Mod(float64(max), 9.0)) + 1, line)
		sum += (int(math.Mod(float64(min), 9.0)) + 1) * 10
		sum += int(math.Mod(float64(max), 9.0)) + 1
		*/

		var digits []int
		for i := 0; i < len(line); i++ {
			d := line[i:]
			if d[0] >= '0' && d[0] <= '9' {
				digits = append(digits, int(d[0] - '0'))
				continue
			}

			for i, num := range(nums) {
				if strings.HasPrefix(d, num) {
					digits = append(digits, i + 1)
					break
				}
			}
		}

		fmt.Printf("min %2d, max %2d, line %s\n", digits[0], digits[len(digits) - 1], line)
		sum += digits[0] * 10 + digits[len(digits) - 1]
	}

	fmt.Println(sum)
}
