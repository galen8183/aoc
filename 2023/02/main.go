package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
	"strconv"
	"strings"
)

type Games [100][][3]int

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

const (
	red = 12
	green = 13
	blue = 14
)

func part1(scanner bufio.Scanner) {
	var sum int
	var idx int

OUTER:
	for scanner.Scan() {
		line := scanner.Text()
		idx++

		for _, draw := range(strings.Split(line, ";")) {
			colours := strings.Split(draw, " ")

			for j, colour := range(colours) {
				var err error
				var r, g, b int

				switch colour {
				case "red", "red,":
					r, err = strconv.Atoi(colours[j - 1])
				case "green", "green,":
					g, err = strconv.Atoi(colours[j - 1])
				case "blue", "blue,":
					b, err = strconv.Atoi(colours[j - 1])
				default:
					// ok
				}

				if err != nil {
					log.Fatal(err)
				}

				if r > red || g > green || b > blue {
					continue OUTER
				}
			}
		}

		sum += idx
	}

	fmt.Println(sum)
}

func part2(scanner bufio.Scanner) {
	var sum float64

	for scanner.Scan() {
		line := scanner.Text()
		var r, g, b float64

		for _, draw := range(strings.Split(line, ";")) {
			colours := strings.Split(draw, " ")

			for j, colour := range(colours) {
				var err error
				var rr, gg, bb int

				switch colour {
				case "red", "red,":
					rr, err = strconv.Atoi(colours[j - 1])
				case "green", "green,":
					gg, err = strconv.Atoi(colours[j - 1])
				case "blue", "blue,":
					bb, err = strconv.Atoi(colours[j - 1])
				default:
					// ok
				}

				if err != nil {
					log.Fatal(err)
				}

				r = math.Max(float64(rr), r)
				g = math.Max(float64(gg), g)
				b = math.Max(float64(bb), b)
			}
		}

		sum += r * g * b
	}

	fmt.Println(sum)
}
