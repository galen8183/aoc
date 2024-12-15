package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("test")
	//input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	d := parse(input)
	part1(d)
	part2(d)
}

func part1(d []File) {
	dd := expand(d)
	last := len(dd) - 1
	for i := 0; i < last; i++ {
		f := dd[i]
		if f != -1 {
			continue
		}

		for j := last; j > i; j-- {
			ff := dd[j]
			if ff == -1 {
				continue
			}
			dd[i] = ff
			dd[j] = f
			last = j
			break
		}
	}

	fmt.Println(checksum(dd))
}

func part2(d []File) {
	for i := len(d) - 1; i > 0; i-- {
		f := &d[i]
		if f.ID == -1 {
			continue
		}

		for j := 0; j < i; j++ {
			ff := d[j]
			if ff.ID != -1 || ff.Size < f.Size {
				continue
			}

			d[j] = *f
			ff.Size -= f.Size
			zero(d, i)

			if ff.Size == 0 {
				break
			}
			fs := make([]File, len(d) + 1)
			copy(fs, d)
			fs = append(fs[:j+1], ff)
			d = append(fs, d[j+1:]...)
			i++
			break
		}
	}

	cs := checksum(expand(d))
	fmt.Println(cs)
}

type File struct {
	Size, ID int
}

func parse(in []byte) []File {
	var fs []File
	for i, n := range(in) {
		if int(n - '0') == 0 {
			continue
		}

		if i != 0 && i / 2 == (i - 1) / 2 {
			i = -2
		}

		fs = append(fs, File{int(n - '0'), i / 2})
	}
	return fs
}

func expand(d []File) (dd []int) {
	for _, f := range(d) {
		for i := 0; i < f.Size; i++ {
			dd = append(dd, f.ID)
		}
	}
	return
}

func zero(d []File, i int) {
	d[i].ID = -1
	if i > 0 && d[i-1].ID == -1 {
		d[i-1].Size += d[i].Size
		d = append(d[:i], d[i+1:]...)
		i--
	}
	if i < len(d) - 1 && d[i+1].ID == -1 {
		d[i].Size += d[i+1].Size
		d = append(d[:i+1], d[i+2:]...)
	}
	return
}

func checksum(d []int) (sum int) {
	for i, f := range(d) {
		if f == -1 {
			continue
		}
		sum += i * f
	}
	return
}
