package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	//part1(m)
	part2(m)
}

type Map struct {
	Map  [][]Cell
	Ants map[byte][]*Cell
}

type Cell struct {
	X, Y  int
	Freq  byte
	Anode bool
}

func parse(sc bufio.Scanner) (m Map) {
	var x int
	m.Ants = make(map[byte][]*Cell)
	for sc.Scan() {
		m.Map = append(m.Map, []Cell{})
		for y, c := range(sc.Bytes()) {
			m.Map[x] = append(m.Map[x], Cell{x, y, 0, false})
			if c != '.' {
				m.Map[x][y].Freq = c
				if m.Ants[c] == nil {
					m.Ants[c] = []*Cell{}
				}
				m.Ants[c] = append(m.Ants[c], &m.Map[x][y])
			}
		}
		x++
	}

	return
}

func part1(m Map) {
	var sum int
	for _, cells := range(m.Ants) {
		for _, c := range(cells) {
			for _, cc := range(cells) {
				if c.X == cc.X && c.Y == cc.Y {
					continue
				}

				distX := int(math.Abs(float64(c.X - cc.X)))
				distY := int(math.Abs(float64(c.Y - cc.Y)))
				if c.X < cc.X {
					distX *= -1
				}
				if c.Y < cc.Y {
					distY *= -1
				}

				if m.Check(c.X + distX, c.Y + distY) {
					sum++
				}
			}
		}
	}

	fmt.Println(m.Print())
	fmt.Println(sum)
}

func part2(m Map) {
	var sum int
	for _, cells := range(m.Ants) {
		for _, c := range(cells) {
			for _, cc := range(cells) {
				if c.X == cc.X && c.Y == cc.Y {
					continue
				}

				distX := int(math.Abs(float64(c.X - cc.X)))
				distY := int(math.Abs(float64(c.Y - cc.Y)))
				if c.X < cc.X {
					distX *= -1
				}
				if c.Y < cc.Y {
					distY *= -1
				}

				x := c.X
				y := c.Y
				for {
					if _, ok := m.GetInBounds(x, y); !ok {
						break
					}

					if m.Check(x, y) {
						sum++
					}
					x += distX
					y += distY
				}
			}
		}
	}

	fmt.Println(m.Print())
	fmt.Println(sum)
}

func (m *Map) Check(x, y int) bool {
	cell, ok := m.GetInBounds(x, y)
	if !ok {
		return false
	}
	if cell.Anode {
		return false
	}
	cell.Anode = true
	return true
}

func (m *Map) GetInBounds(x, y int) (*Cell, bool) {
	if x < 0 || x >= len(m.Map) {
		return nil, false
	}
	if y < 0 || y >= len(m.Map[x]) {
		return nil, false
	}
	return &m.Map[x][y], true
}

func (m *Map) Print() (buf string) {
	for _, c := range(m.Map) {
		for _, cell := range(c) {
			c := '.'
			if cell.Freq != 0 {
				c = rune(cell.Freq)
			} else if cell.Anode {
				c = '#'
			}
			buf = fmt.Sprintf("%s%c", buf, c)
		}
		buf = fmt.Sprint(buf, "\n")
	}
	return
}
