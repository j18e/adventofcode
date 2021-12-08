package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	Tree  rune = '#'
	Empty rune = '.'
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	for _, l := range strings.Split(string(bs), "\n") {
		if l != "" {
			lines = append(lines, l)
		}
	}

	grid := &Grid{lines: lines}
	routes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	res := 1
	for _, r := range routes {
		grid.Reset()
		trees := grid.CountTrees(r[0], r[1])
		fmt.Println(trees)
		res *= trees
	}
	fmt.Println(res)
}

type Grid struct {
	lines      []string
	posX, posY int
}

func (g *Grid) Reset() {
	g.posX = 0
	g.posY = 0
}

func (g *Grid) CountTrees(right, down int) int {
	trees := 0
	for g.On() {
		if g.Current() == Tree {
			trees++
		}
		g.Right(right)
		g.Down(down)
	}
	return trees
}

// On reports whether the position has gone beyond the grid.
func (g *Grid) On() bool {
	return g.posY < len(g.lines)
}

func (g *Grid) Current() rune {
	return rune(g.lines[g.posY][g.posX])
}

func (g *Grid) Right(i int) {
	res := g.posX + i
	if l := len(g.lines[0]); res >= l {
		res -= l
	}
	g.posX = res
}

func (g *Grid) Down(i int) {
	g.posY += i
}
