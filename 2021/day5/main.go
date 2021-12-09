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
	fmt.Println(part2(lines))
}

type Coord struct {
	X, Y int
}

type Line struct {
	Start Coord
	End   Coord
}

func (l Line) Diagonal() bool {
	return !(l.Start.X == l.End.X || l.Start.Y == l.End.Y)
}

func (l Line) Points() []Coord {
	x, y := l.Start.X, l.Start.Y
	if x == l.End.X && y == l.End.Y {
		return []Coord{{x, y}}
	}
	res := []Coord{{x, y}}
	for {
		if x < l.End.X {
			x++
		} else if x > l.End.X {
			x--
		}
		if y < l.End.Y {
			y++
		} else if y > l.End.Y {
			y--
		}
		res = append(res, Coord{x, y})
		if x == l.End.X && y == l.End.Y {
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
			Start: Coord{converting.Atoi(matches[1]), converting.Atoi(matches[2])},
			End:   Coord{converting.Atoi(matches[3]), converting.Atoi(matches[4])},
		})
	}
	return res
}

func part1(lines []Line) int {
	return countOverlaps(lines, true)
}

func part2(lines []Line) int {
	return countOverlaps(lines, false)
}

func countOverlaps(lines []Line, skipDiagonal bool) int {
	counts := make(map[Coord]int)
	for _, line := range lines {
		if skipDiagonal && line.Diagonal() {
			continue
		}
		for _, point := range line.Points() {
			counts[point]++
		}
	}
	var res int
	for _, cnt := range counts {
		if cnt > 1 {
			res++
		}
	}
	return res
}
