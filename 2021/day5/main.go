package main

import (
	"fmt"
	"regexp"

	"github.com/j18e/adventofcode/pkg/converting"
	"github.com/j18e/adventofcode/pkg/inputting"
)

var reInputLine = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func main() {
	lines := parseInput(inputting.GetInputStrings("input.txt"))
	fmt.Println(part1(lines))
	// fmt.Println(part2(input))
}

type Coordinate struct {
	X, Y int
}

type Line struct {
	X1, X2, Y1, Y2 int
}

func (l Line) Diagonal() bool {
	return l.X1 == l.X2 || l.Y1 == l.Y2
}

func (l Line) Points() []Coordinate {
	var res []Coordinate
	var x, y int
	for {
		if x < l.X2 {
			x++
		}
		if y < l.Y2 {
			y++
		}
		res = append(res, Coordinate{x, y})
		if x == l.X2 && y == l.Y2 {
			break
		}
	}
	return res
}

func parseInput(input []string) []Line {
	var res []Line
	for _, ln := range input {
		matches := reInputLine.FindStringSubmatch(ln)
		res = append(res, Line{
			X1: converting.Atoi(matches[1]),
			X2: converting.Atoi(matches[2]),
			Y1: converting.Atoi(matches[3]),
			Y2: converting.Atoi(matches[4]),
		})
	}
	return res
}

func part1(lines []Line) int {
	counts := make(map[Coordinate]int)
	for _, line := range lines {
		if line.Diagonal() {
			continue
		}
	}

	return 0
}

func part2(input string) int {
	return 0
}
