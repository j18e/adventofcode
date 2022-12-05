package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bs, err := os.ReadFile("input.txt")
	if err != nil {
		return err
	}
	lines := strings.Split(string(bs), "\n")
	var pairs []Pair
	for _, ln := range lines {
		if ln == "" {
			continue
		}
		var pair Pair
		split := strings.Split(ln, ",")
		elf1 := strings.Split(split[0], "-")
		elf1Start, err := strconv.Atoi(elf1[0])
		check(err)
		elf1End, err := strconv.Atoi(elf1[1])
		check(err)
		for i := elf1Start; i <= elf1End; i++ {
			pair.Elf1 = append(pair.Elf1, i)
		}
		elf2 := strings.Split(split[1], "-")
		elf2Start, err := strconv.Atoi(elf2[0])
		check(err)
		elf2End, err := strconv.Atoi(elf2[1])
		check(err)
		for i := elf2Start; i <= elf2End; i++ {
			pair.Elf2 = append(pair.Elf2, i)
		}
		pairs = append(pairs, pair)
	}

	total1, total2 := 0, 0
	for _, p := range pairs {
		if p.FullOverlap() {
			total1++
		}
		if p.AnyOverlap() {
			total2++
		}
	}
	fmt.Println(total1, total2)
	return nil
}

type Pair struct {
	Elf1, Elf2 []int
}

func (p Pair) FullOverlap() bool {
	if p.Elf1Start() >= p.Elf2Start() && p.Elf1End() <= p.Elf2End() {
		return true
	}
	if p.Elf2Start() >= p.Elf1Start() && p.Elf2End() <= p.Elf1End() {
		return true
	}
	return false
}

func (p Pair) AnyOverlap() bool {
	if p.Elf1Start() >= p.Elf2Start() && p.Elf1Start() <= p.Elf2End() {
		return true
	}
	if p.Elf2Start() >= p.Elf1Start() && p.Elf2Start() <= p.Elf1End() {
		return true
	}
	return false
}

func (p Pair) Elf1Start() int {
	return p.Elf1[0]
}

func (p Pair) Elf1End() int {
	return p.Elf1[len(p.Elf1)-1]
}

func (p Pair) Elf2Start() int {
	return p.Elf2[0]
}

func (p Pair) Elf2End() int {
	return p.Elf2[len(p.Elf2)-1]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
