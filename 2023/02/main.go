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

func main() {
	// input, err := os.Open("test")
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	games := parse(*scanner)
	part1(games)
	part2(games)
}

const (
	red = 12
	green = 13
	blue = 14
)

func part1(games Games) {
	var sum int

OUTER:
	for i, game := range(games) {
		for _, draw := range(game) {
			if draw.Red > red || draw.Green > green || draw.Blue > blue {
				continue OUTER
			}
		}

		sum += i + 1
	}

	fmt.Println(sum)
}

func part2(games Games) {
	var sum float64

	for _, game := range(games) {
		var r, g, b float64
		for _, draw := range(game) {
			r = math.Max(float64(draw.Red), r)
			g = math.Max(float64(draw.Green), g)
			b = math.Max(float64(draw.Blue), b)
		}

		sum += r * g * b
	}

	fmt.Println(sum)
}

type Draw struct {
	Red   int
	Green int
	Blue  int
}

type Games [100][]Draw

func parse(scanner bufio.Scanner) Games {
	var out Games
	var idx int

	for scanner.Scan() {
		line := scanner.Text()
		var game []Draw

		for _, draw := range(strings.Split(line, ";")) {
			colours := strings.Split(draw, " ")
			var r, g, b int

			for i, colour := range(colours) {
				var err error

				switch colour {
				case "red", "red,":
					r, err = strconv.Atoi(colours[i - 1])
				case "green", "green,":
					g, err = strconv.Atoi(colours[i - 1])
				case "blue", "blue,":
					b, err = strconv.Atoi(colours[i - 1])
				default:
					// ok
				}

				if err != nil {
					log.Fatal(err)
				}
			}

			game = append(game, Draw{r, g, b})
		}

		out[idx] = game
		idx++
	}

	return out
}
