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
	solve(*scanner, false)
	//solve(*scanner, true)
}

func solve(sc bufio.Scanner, p2 bool) {
	var sum int
	for sc.Scan() {
		line := strings.Split(sc.Text(), ": ")
		var operands []int
		for _, n := range(strings.Split(line[1], " ")) {
			m, _ := strconv.Atoi(n)
			operands = append(operands, m)
		}

		res, _ := strconv.Atoi(line[0])
		if check(res, operands, p2) {
			sum += res
		}
	}

	fmt.Println(sum)
}

func check(res int, op []int, p2 bool) bool {
	if len(op) == 1 {
		return res == op[0]
	}

	lastIdx := len(op) - 1
	last := op[lastIdx]
	var canAdd, canMul, canConc bool

	canAdd = check(res - last, op[:lastIdx], p2)
	if res % last == 0 {
		canMul = check(res / last, op[:lastIdx], p2)
	}
	if p2 {
		pow := 1
		for pow <= last {
			pow *= 10
		}

		if (res - last) % pow == 0 {
			canConc = check(res / pow, op[:lastIdx], p2)
		}
	}

	return canMul || canAdd || canConc
}
