package main

import "fmt"

type Spot int

const (
	SpotFloor Spot = iota
	SpotEmpty
	SpotOccupied
	SpotOffGrid
)

func ParseSpot(r rune) Spot {
	switch r {
	case '.':
		return SpotFloor
	case 'L':
		return SpotEmpty
	case '#':
		return SpotOccupied
	}
	panic(fmt.Sprintf("parsing %s: must be '.', 'L' or '#'", string(r)))
}
