package main

import "fmt"

type Direction int

const (
	N Direction = iota
	S
	E
	W
	L
	R
	F
)

func (d Direction) String() string {
	var dir string
	switch d {
	case N:
		dir = "North"
	case S:
		dir = "South"
	case E:
		dir = "East"
	case W:
		dir = "West"
	case L:
		dir = "Left"
	case R:
		dir = "Right"
	case F:
		dir = "Forward"
	}
	return dir
}

type Instruction struct {
	D Direction
	N int
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s %d", i.D, i.N)
}
