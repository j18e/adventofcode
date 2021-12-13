package main

import (
	"fmt"
	"strings"

	"github.com/j18e/adventofcode/pkg/converting"
	"github.com/j18e/adventofcode/pkg/inputting"
)

type Coord struct {
	x, y int
}

type Octopus struct {
	energy  int
	flashed bool
}

var directions = []Coord{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func main() {
	grid := parseInput(inputting.GetInputStrings("input.txt"))
	fmt.Println(part1(grid))
	fmt.Println(part2(grid))
}

func parseInput(input []string) Grid {
	var grid Grid
	for _, ln := range input {
		var row []*Octopus
		for _, c := range ln {
			o := Octopus{converting.Atoi(string(c)), false}
			row = append(row, &o)
		}
		grid = append(grid, row)
	}
	return grid
}

func part1(grid Grid) int {
	return grid.StepN(100)
}

func part2(grid Grid) int {
	size := grid.Size()
	for i := 1; ; i++ {
		flashes := grid.Step()
		if flashes == size {
			return i
		}
	}
}

type Grid [][]*Octopus

func (g Grid) Size() int {
	return len(g) * len(g[0])
}

func (g Grid) StepN(n int) int {
	var flashes int
	for i := 0; i < n; i++ {
		flashes += g.Step()
	}
	return flashes
}

func (g Grid) String() string {
	var res string
	for y := range g {
		res += "\n"
		for x := range g[y] {
			res += fmt.Sprintf("%d", g[y][x].energy)
		}
	}
	return strings.TrimSpace(res)
}

func (g Grid) Items() []Point {
	var res []Point
	for y := range g {
		for x := range g[y] {
			res = append(res, Point{x, y, g[y][x]})
		}
	}
	return res
}

type Point struct {
	x, y int
	oct  *Octopus
}

func (g Grid) Step() int {
	for _, p := range g.Items() {
		if !p.oct.flashed {
			p.oct.energy++
			if p.oct.energy > 9 {
				g.flash(p.x, p.y)
			}
		}
	}
	flashed := g.flashing()
	g.resetFlashing()
	return flashed
}

func (g Grid) flash(x, y int) {
	oct := g[y][x]
	oct.flashed = true
	oct.energy = 0
	for _, d := range directions {
		nx, ny := x+d.x, y+d.y
		if nx < 0 || nx >= len(g[0]) || ny < 0 || ny >= len(g) {
			continue
		}
		noct := g[ny][nx]
		if !noct.flashed {
			noct.energy++
			if noct.energy > 9 {
				g.flash(nx, ny)
			}
		}
	}
}

func (g Grid) flashing() int {
	var res int
	for _, p := range g.Items() {
		if p.oct.flashed {
			res++
		}
	}
	return res
}

func (g Grid) resetFlashing() {
	for _, p := range g.Items() {
		p.oct.flashed = false
	}
}
