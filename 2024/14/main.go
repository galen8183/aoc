package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//input, err := os.Open("test")
	//w, h := 11, 7
	input, err := os.Open("input")
	w, h := 101, 103
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	l := NewLobby(w, h)
	sc := bufio.NewScanner(input)
	l.parse(sc)
	//part1(&l)
	part2(&l)
}

func part1(l *Lobby) {
	l.stepAll(100)

	quads := make(map[bool]map[bool]int)
	quads[true] = make(map[bool]int)
	quads[false] = make(map[bool]int)
	for _, r := range(l.R) {
		// skip middle column || row
		if r.X == l.Width / 2 || r.Y == l.Height / 2 {
			continue
		}
		quads[r.X < l.Width / 2][r.Y < l.Height / 2]++
	}

	sum := 1
	for _, q := range(quads) {
		for _, qq := range(q) {
			sum *= qq
		}
	}

	fmt.Println(sum)
}

func part2(l *Lobby) {
	var cont byte
	i := 0
	for {
		if l.maybeHasTree() {
			fmt.Println(l, "at step", i)
			fmt.Print("continue [y/N]? ")
			fmt.Scanf("%c", &cont)
			if cont != 'y' {
				break
			}
		}
		l.stepAll(1)
		i++
	}
}

type Lobby struct {
	Width, Height int
	Map [][]Tile
	R []*Robot
}

type Tile struct {
	X, Y int
	R map[*Robot]bool
}

type Robot struct {
	X, Y int
	Vel [2]int
}

func NewLobby(w, h int) Lobby {
	l := Lobby{w, h, make([][]Tile, h), []*Robot{}}
	for i := 0; i < h; i++ {
		l.Map[i] = make([]Tile, w)
		for j := 0; j < w; j++ {
			l.Map[i][j].R = make(map[*Robot]bool)
		}
	}
	return l
}

func (l *Lobby) parse(sc *bufio.Scanner) {
	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")
		p := strings.Split(strings.Split(line[0], "=")[1], ",")
		v := strings.Split(strings.Split(line[1], "=")[1], ",")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])

		r := &Robot{x, y, [2]int{vx, vy}}
		l.Map[y][x].R[r] = true
		l.R = append(l.R, r)
	}
	return
}

func (l *Lobby) stepAll(steps int) {
	for _, r := range(l.R) {
		l.step(r, steps)
	}
	return
}

func (l *Lobby) step(r *Robot, steps int) {
	x := r.X + (r.Vel[0] * steps)
	x = int(math.Mod(float64(x), float64(l.Width)))
	if x < 0 {
		x += l.Width
	} else if x >= l.Width {
		x = l.Width
	}

	y := r.Y + (r.Vel[1] * steps)
	y = int(math.Mod(float64(y), float64(l.Height)))
	if y < 0 {
		y += l.Height
	} else if y >= l.Height {
		y -= l.Height
	}

	delete(l.Map[r.Y][r.X].R, r)
	r.X, r.Y = x, y
	l.Map[y][x].R[r] = true
}

func (l *Lobby) String() (s string) {
	for _, row := range(l.Map) {
		for _, t := range(row) {
			if len(t.R) != 0 {
				s += strconv.Itoa(len(t.R))
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return
}

func (l *Lobby) maybeHasTree() bool {
OUTER:
	for i := 0; i < l.Height; i++ {
		var count, last int
		for j, t := range(l.Map[i]) {
			if len(t.R) == 0 {
				continue
			}
			if last != 0 && j > last + 2 {
				continue OUTER
			}
			count++
			if count > 15 {
				return true
			}
			last = j
		}
	}

	return false
}
