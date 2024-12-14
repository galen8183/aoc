package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("test")
	//input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	solve(string(input), 0)
	solve(string(input), 10000000000000)
}

func solve(in string, offset int) {
	var sum int

	for _, m := range(strings.Split(in, "\n\n")) {
		mm := strings.Split(m, "\n")
		a_x, a_y := ints(mm[0])
		b_x, b_y := ints(mm[1])
		p_x, p_y := ints(mm[2])
		p_x += offset
		p_y += offset

		dA := a_x * b_y - a_y * b_x
		if dA == 0 {
			continue
		}

		a := (b_y * p_x - b_x * p_y) / dA
		b := (a_x * p_y - a_y * p_x) / dA
		if a * a_x + b * b_x == p_x && a * a_y + b * b_y == p_y {
			sum += a * 3 + b
		}
	}

	fmt.Println(sum)
}

func ints(in string) (int, int) {
	line := strings.Split(strings.Split(in, ": ")[1], ", ")
	a, _ := strconv.Atoi(line[0][2:])
	b, _ := strconv.Atoi(line[1][2:])
	return a, b
}
