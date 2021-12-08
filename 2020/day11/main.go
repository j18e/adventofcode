package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	part1(loadGrid())
	part2(loadGrid())
}

func part1(grid *Grid) {
	grid.FillSeats()
	var prev, cnt int
	for {
		grid.FillSeats()
		cur := grid.TotalOccupied()
		if cur == prev {
			break
		}
		prev = cur
		cnt++
	}
	fmt.Println(prev, cnt)
}

func part2(grid *Grid) {
	grid.FillSeats2()
	var prev, cnt int
	for {
		grid.FillSeats2()
		cur := grid.TotalOccupied()
		if cur == prev {
			break
		}
		prev = cur
		cnt++
	}
	fmt.Println(prev, cnt)
}

func loadGrid() *Grid {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// assemble grid
	var data []string
	for _, l := range strings.Split(string(bs), "\n") {
		if l == "" {
			continue
		}
		data = append(data, l)
	}
	grid := &Grid{len(data), len(data[0]), make(map[int]Spot)}
	for row := range data {
		for seat, r := range data[row] {
			grid.Spots[grid.SeatID(coordinate{row, seat})] = ParseSpot(r)
		}
	}
	return grid
}
