package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("test")
	//input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	m := parse(*scanner)
	part1(m)
	part2(m)
}

type Map struct {
	Map   [][]Pos
	Heads []Pos
}

type Pos struct {
	X, Y  int
	Elev  int
}

func parse(sc bufio.Scanner) Map {
	var m Map
	var i int
	for sc.Scan() {
		var row []Pos
		for j, n := range(sc.Bytes()) {
			nn := int(n - '0')
			row = append(row, Pos{i, j, nn})
			if nn == 0 {
				m.Heads = append(m.Heads, row[j])
			}
		}
		m.Map = append(m.Map, row)
		i++
	}
	return m
}

func part1(m Map) {
	var sum int
	for _, h := range(m.Heads) {
		peaks := m.step(h)
		var counted []Pos
		for _, p := range(peaks) {
			if !contains(counted, p) {
				sum++
				counted = append(counted, p)
			}
		}
	}

	fmt.Println(sum)
}

func part2(m Map) {
	var sum int
	for _, h := range(m.Heads) {
		peaks := m.step(h)
		sum += len(peaks)
	}
	fmt.Println(sum)
}

func (m *Map) step(p Pos) (peaks []Pos) {
	var dirs = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range(dirs) {
		pp, ok := m.getInBounds(p.X + d[0], p.Y + d[1])
		if !ok {
			continue
		}
		if pp.Elev != p.Elev+1 {
			continue
		}
		if pp.Elev == 9 {
			peaks = append(peaks, pp)
		} else {
			peaks = append(peaks, m.step(pp)...)
		}
	}
	return
}

func (m *Map) getInBounds(x, y int) (Pos, bool) {
	if x < 0 || y < 0 || x >= len(m.Map) || y >= len(m.Map[x]) {
		return Pos{}, false
	}
	return m.Map[x][y], true
}

func contains(l []Pos, p Pos) bool {
	for _, pp := range(l) {
		if p.X == pp.X && p.Y == pp.Y {
			return true
		}
	}
	return false
}
