package main

import (
	"fmt"
	"sort"

	"github.com/j18e/adventofcode/pkg/converting"
	"github.com/j18e/adventofcode/pkg/inputting"
)

func main() {
	input := parseInput(inputting.GetInputStrings("input.txt"))
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func parseInput(input []string) Grid {
	grid := make(Grid, len(input))
	for i, ln := range input {
		for _, n := range ln {
			grid[i] = append(grid[i], converting.Atoi(string(n)))
		}
	}
	return grid
}

func part1(grid Grid) int {
	var res int
	for _, coord := range grid.lowPoints() {
		risk := grid.Coord(coord) + 1
		res += risk
	}
	return res
}

func part2(grid Grid) int {
	var bs []int
	for _, c := range grid.lowPoints() {
		bs = append(bs, grid.BasinSize(c))
	}
	sort.Ints(bs)
	bs = bs[len(bs)-4:]
	return bs[len(bs)-1] * bs[len(bs)-2] * bs[len(bs)-3]
}

type Coord struct {
	x, y int
}

type Grid [][]int

func (g Grid) BasinSize(c Coord) int {
	seen := map[Coord]bool{
		c: true,
	}
	queue := []Coord{c}
	var current Coord
	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		for _, next := range g.AdjacentUnder9(current) {
			if seen[next] {
				continue
			}
			queue = append(queue, next)
			seen[next] = true
		}
	}
	return len(seen)
}

func (g Grid) AdjacentUnder9(c Coord) []Coord {
	var res []Coord
	for _, cc := range g.AdjacentCoords(c) {
		if g.Coord(cc) == 9 {
			continue
		}
		res = append(res, cc)
	}
	return res
}

func coordInList(c Coord, cx []Coord) bool {
	for _, cc := range cx {
		if c.x == cc.x && c.y == cc.y {
			return true
		}
	}
	return false
}

func (g Grid) AdjacentCoords(c Coord) []Coord {
	var res []Coord
	x, y := c.x, c.y
	if x > 0 {
		res = append(res, Coord{x - 1, y})
	}
	if x < len(g[0])-1 {
		res = append(res, Coord{x + 1, y})
	}
	if y > 0 {
		res = append(res, Coord{x, y - 1})
	}
	if y < len(g)-1 {
		res = append(res, Coord{x, y + 1})
	}
	return res
}

func (g Grid) Coords() []Coord {
	var res []Coord
	for y := range g {
		for x := range g[y] {
			res = append(res, Coord{x, y})
		}
	}
	return res
}

func (g Grid) Coord(c Coord) int {
	return g[c.y][c.x]
}

func (g Grid) CoordXY(x, y int) int {
	return g[y][x]
}

func (g Grid) lowPoints() []Coord {
	var res []Coord
	for _, c := range g.Coords() {
		adj := g.AdjacentCoords(c)
		if isLeast(c, adj, g) {
			res = append(res, c)
		}
	}
	return res
}

func isLeast(c Coord, cx []Coord, grid Grid) bool {
	num := grid.Coord(c)
	for _, cc := range cx {
		if num >= grid.Coord(cc) {
			return false
		}
	}
	return true
}
