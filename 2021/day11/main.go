package main

import (
	"fmt"
	"strings"

	"github.com/j18e/adventofcode/pkg/converting"
	"github.com/j18e/adventofcode/pkg/inputting"
)

func main() {
	grid := parseInput(inputting.GetInputStrings("input.txt"))
	fmt.Println(part1(grid))
	fmt.Println(part2(grid))
}

func parseInput(input []string) Grid {
	var grid Grid
	for _, ln := range input {
		grid = append(grid, func(s string) []int {
			var ix []int
			for i := range s {
				ix = append(ix, converting.Atoi(string(s[i])))
			}
			return ix
		}(ln))
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
		fmt.Println(flashes)
	}
}

type GridItem struct {
	x, y, val int
}

func (i GridItem) Flashing() bool {
	return i.val > 9
}

type Grid [][]int

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
			res += fmt.Sprintf("%d", g[y][x])
		}
	}
	return strings.TrimSpace(res)
}

func (g Grid) Items() []GridItem {
	var res []GridItem
	for y := range g {
		for x := range g[y] {
			res = append(res, GridItem{x, y, g[y][x]})
		}
	}
	return res
}

func (g Grid) Apply(item GridItem) {
	g[item.y][item.x] = item.val
}

func (g Grid) Step() int {
	flashed := make(map[[2]int]bool)
	var queue []GridItem
	for _, item := range g.Items() {
		item.val++
		g.Apply(item)
		if item.Flashing() {
			flashed[[2]int{item.x, item.y}] = true
			queue = append(queue, item)
		}
	}
	var current GridItem
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		for _, item := range g.adjacent(current) {
			if flashed[[2]int{item.x, item.y}] {
				continue
			}
			item.val++
			g.Apply(item)
			if item.Flashing() {
				flashed[[2]int{item.x, item.y}] = true
				queue = append(queue, item)
			}
		}
	}
	g.resetFlashing()
	return len(flashed)
}

func (g Grid) adjacent(i GridItem) []GridItem {
	var res []GridItem
	x, y := i.x, i.y

	// topleft
	if x > 0 && y > 0 {
		res = append(res, GridItem{x - 1, y - 1, g[y-1][x-1]})
	}
	// top
	if y > 0 {
		res = append(res, GridItem{x, y - 1, g[y-1][x]})
	}
	// topright
	if x < len(g[0])-1 && y > 0 {
		res = append(res, GridItem{x + 1, y - 1, g[y-1][x+1]})
	}
	// left
	if x > 0 {
		res = append(res, GridItem{x - 1, y, g[y][x-1]})
	}
	// right
	if x < len(g[0])-1 {
		res = append(res, GridItem{x + 1, y, g[y][x+1]})
	}
	// bottomleft
	if x > 0 && y < len(g)-1 {
		res = append(res, GridItem{x - 1, y + 1, g[y+1][x-1]})
	}
	// bottom
	if y < len(g)-1 {
		res = append(res, GridItem{x, y + 1, g[y+1][x]})
	}
	// bottomright
	if x < len(g[0])-1 && y < len(g)-1 {
		res = append(res, GridItem{x + 1, y + 1, g[y+1][x+1]})
	}
	return res
}

func (g Grid) resetFlashing() {
	for _, item := range g.Items() {
		if item.val > 9 {
			item.val = 0
			g.Apply(item)
		}
	}
}
