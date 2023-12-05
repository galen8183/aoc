package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	cards := parse(*scanner)

	// part1(cards)
	part2(cards)
}

type Card struct {
	CardNum int
	WinNums []int
	OwnNums []int
	Wins    int
	Count   int
}

func part1(cards []Card) {
	var sum int
	for _, card := range(cards) {
		var wins int
		for _, i := range(card.OwnNums) {
			for _, j := range(card.WinNums) {
				if i == j {
					wins = max(wins * 2, 1)
					break
				}
			}
		}

		sum += wins
	}

	fmt.Println(sum)
}

func part2(cards []Card) {
	var sum int

	for i := 0; i < len(cards); i++ {
		card := &cards[i]
		for _, n := range(card.OwnNums) {
			for _, m := range(card.WinNums) {
				if n == m {
					card.Wins++
					break
				}
			}
		}

		for j := 0; j < card.Count; j++ {
			sum++
			for k := 0; k < card.Wins; k++ {
				cards[i + k + 1].Count++
			}
		}
	}

	fmt.Println(sum)
}

func parse(scanner bufio.Scanner) (cards []Card) {
	var cardNum int

	for scanner.Scan() {
		var win, own bool
		card := Card{CardNum: cardNum, Count: 1}
		cardNum++

		for _, field := range(strings.Split(scanner.Text(), " ")) {
			switch field {
			case "|":
				win = false
				own = true
				continue
			case "", "Card":
				continue
			default:
				if field[len(field) - 1] == ':' {
					win = true
					continue
				}
			}

			n, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}

			if win {
				card.WinNums = append(card.WinNums, n)
			}
			if own {
				card.OwnNums = append(card.OwnNums, n)
			}
		}

		cards = append(cards, card)
	}

	return
}
