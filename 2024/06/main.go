package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var animate bool
func init() {
	if len(os.Args) <= 1 {
		return
	}
	animate = os.Args[len(os.Args)-1] == "animate"
}

func main() {
	//input, err := os.Open("test")
	input, err := os.Open("input")
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
	Map     [][]Cell
	Touched map[int]map[int]int
	Guard   Guard
}

type Cell struct {
	X, Y    int
	Obst    bool
	Touched bool
}

type Guard struct {
	X, Y int
	Rot  int
}

const (
	Rot1 = 1
	Rot2 = 1 << 1
	Rot3 = 1 << 2
	Rot4 = 1 << 3
)

func parse(scanner bufio.Scanner) (m Map) {
	var x int
	for scanner.Scan() {
		m.Map = append(m.Map, []Cell{})
		for y, c := range(scanner.Bytes()) {
			m.Map[x] = append(m.Map[x], Cell{x, y, c == '#', false})
			if c == '^' {
				m.Guard = Guard{x, y, 1}
				m.Map[x][y].Touched = true
				m.Touched = make(map[int]map[int]int)
				m.Touched[x] = make(map[int]int)
				m.Touched[x][y] |= m.Guard.Rot
			}
		}
		x++
	}

	return
}

func (m *Map) Walk() (bool, bool) {
	x := m.Guard.X
	y := m.Guard.Y
	if m.Guard.Rot & Rot1 != 0 {
		x--
	} else if m.Guard.Rot & Rot2 != 0 {
		y++
	} else if m.Guard.Rot & Rot3 != 0 {
		x++
	} else if m.Guard.Rot & Rot4 != 0 {
		y--
	}

	if x < 0 || x >= len(m.Map) || y < 0 || y >= len(m.Map[x]) {
		// reached the edge, we're done
		return true, false
	}

	c := m.Map[x][y]
	if c.Obst {
		m.Guard.Rotate()
		return false, false
	}

	m.Guard.X = x
	m.Guard.Y = y
	if !m.Map[x][y].Touched {
		m.Map[x][y].Touched = true
		if m.Touched[x] == nil {
			m.Touched[x] = make(map[int]int)
		}
		m.Touched[x][y] |= m.Guard.Rot
	} else if m.Touched[x][y] & m.Guard.Rot != 0 {
		return true, true
	}

	return false, false
}

func (m *Map) Print() (buf string) {
	for _, c := range(m.Map) {
		for _, cell := range(c) {
			c := "."
			if m.Guard.X == cell.X && m.Guard.Y == cell.Y {
				switch m.Guard.Rot {
				case Rot1: c = "^"
				case Rot2: c = ">"
				case Rot3: c = "v"
				case Rot4: c = "<"
				}
			} else if cell.Touched {
				c = "X"
			} else if cell.Obst {
				c = "#"
			}
			buf = fmt.Sprint(buf, c)
		}
		buf = fmt.Sprint(buf, "\n")
	}
	return
}

func (m *Map) Reset() {
	for x, c := range(m.Touched) {
		for y := range(c) {
			m.Map[x][y].Touched = false
		}
	}
	m.Touched = make(map[int]map[int]int)
}

func (g *Guard) Rotate() {
	g.Rot <<= 1
	if g.Rot > 1 << 3 {
		g.Rot = 1
	}
	return
}

func part1(m Map) {
	for {
		done, _ := m.Walk()
		if done {
			break
		}
	}

	var sum int
	for _, t := range(m.Touched) {
		sum += len(t)
	}
	fmt.Println(sum)
}

func part2(m Map) {
	init := m.Guard
	for {
		done, _ := m.Walk()
		if done {
			break
		}
	}

	var sum int
	for x, c := range(m.Touched) {
		for y := range(c) {
			m.Reset()
			m.Guard = init
			m.Map[x][y].Obst = true
			for {
				done, loop := m.Walk()
				if done {
					if loop {
						sum++
					}
					break
				}
			}
			m.Map[x][y].Obst = false
		}
	}

	fmt.Println(sum)
}
