package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"regexp"
)

func main() {
	//input, err := os.ReadFile("test")
	input, err := os.ReadFile("test2")
	//input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	part1(input)
	part2(input)
}

func part1(in []byte) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	insts := re.FindAll(in, -1)
	var sum int
	for _, i := range(insts) {
		ii := strings.Split(string(i[4:len(i)-1]), ",")
		n, _ := strconv.Atoi(ii[0])
		m, _ := strconv.Atoi(ii[1])
		sum += n * m
	}
	fmt.Println(sum)
}

func part2(in []byte) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do(?:n't)?\(\)`)
	insts := re.FindAll(in, -1)
	var sum int
	do := true
	for _, i := range(insts) {
		if string(i) == "do()" {
			do = true
			continue
		}
		if !do || string(i) == "don't()" {
			do = false
			continue
		}

		ii := strings.Split(string(i[4:len(i)-1]), ",")
		n, _ := strconv.Atoi(ii[0])
		m, _ := strconv.Atoi(ii[1])
		sum += n * m
	}
	fmt.Println(sum)
}
