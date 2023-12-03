package main

import (
	"bufio"
	"fmt"
	"log"
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
	part1(*scanner)
	// part2(*scanner)
}

func part1(scanner bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
	}
}

func part2(scanner bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
	}
}
