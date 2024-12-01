package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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
	//part1(*scanner)
	part2(*scanner)
}

func part1(scanner bufio.Scanner) {
	var col1, col2 []int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(line[0])
		m, _ := strconv.Atoi(line[1])

		col1 = append(col1, n)
		col2 = append(col2, m)
	}

	slices.Sort(col1)
	slices.Sort(col2)

	var sum int
	for i := 0; i < len(col1); i++ {
		distance := col1[i] - col2[i]
		sum += int(math.Abs(float64(distance)))
	}

	fmt.Print(sum)
}

func part2(scanner bufio.Scanner) {
	var col1 []int
	col2 := make(map[int]int)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(line[0])
		m, _ := strconv.Atoi(line[1])

		col1 = append(col1, n)
		col2[m]++
	}

	var sim int
	for _, i := range(col1) {
		sim += i * col2[i]
	}

	fmt.Print(sim)
}
