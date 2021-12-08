package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	takenSeats := make([][]bool, 128)
	for r := range takenSeats {
		takenSeats[r] = make([]bool, 8)
	}

	var highest int
	for _, l := range strings.Split(string(bs), "\n") {
		if l == "" {
			continue
		}
		row := findRow(l[0:7])
		seat := findSeat(l[7:])
		takenSeats[row][seat] = true
		if id := getId(row, seat); id > highest {
			highest = id
		}
	}
	fmt.Println("highest id:", highest)

	var mySeat, prev int
	loop:
	for r := range takenSeats {
		for s := range takenSeats[r] {
			if takenSeats[r][s] {
				continue
			}
			id := getId(r, s)
			if id == 0 {
				continue
			}
			if id == prev+1 {
				prev = id
				continue
			}
			mySeat = id
			break loop
		}
	}

	fmt.Println("your seat:", mySeat)
}

func getId(row, seat int) int {
	return row*8+seat
}

func findRow(s string) int {
	var res float64
	for i, r := range s {
		switch r {
		case 'F':
		case 'B':
			res += math.Pow(2, 6-float64(i))
		default:
			panic(fmt.Sprintf("line %s: contains non 'F' or 'B' char", s))
		}
	}
	return int(res)
}

func findSeat(s string) int {
	var res float64
	for i, r := range s {
		switch r {
		case 'L':
		case 'R':
			res += math.Pow(2, 2-float64(i))
		default:
			panic(fmt.Sprintf("line %s: contains non 'L' or 'R' char", s))
		}
	}
	return int(res)
}
