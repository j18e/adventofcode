package main

import (
	"fmt"
	"log"

	"github.com/j18e/adventofcode/2022/common"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := common.ReadInput("input.txt")
	part1 := 0
	part2 := 0
	for y := range input {
		for x := range input[y] {
			if isVisible(input, x, y) {
				part1++
			}
			score := scenicScore(input, x, y)
			if score > part2 {
				part2 = score
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
	return nil
}

func scenicScore(grid []string, x, y int) int {
	if x == 0 || y == 0 || x == len(grid[0])-1 || y == len(grid)-1 {
		return 0
	}
	height := grid[y][x]

	var left, right, up, down int

	// check left
	for i := x - 1; i >= 0; i-- {
		left++
		tree := grid[y][i]
		if tree >= height {
			break
		}
	}
	// check right
	for i := range grid[y][x+1:] {
		right++
		tree := grid[y][i+x+1]
		if tree >= height {
			break
		}
	}
	// check up
	for i := y - 1; i >= 0; i-- {
		up++
		tree := grid[i][x]
		if tree >= height {
			break
		}
	}
	// check down
	for i := range grid[y+1:] {
		down++
		tree := grid[i+y+1][x]
		if tree >= height {
			break
		}
	}
	return left * right * up * down
}

func isVisible(grid []string, x, y int) bool {
	if x == 0 || y == 0 || x == len(grid[0])-1 || y == len(grid)-1 {
		return true
	}
	height := grid[y][x]

	left, right, up, down := true, true, true, true

	// check left
	for i := range grid[y][:x] {
		if grid[y][i] >= height {
			left = false
			break
		}
	}
	if left {
		return true
	}
	// check right
	for i := range grid[y][x+1:] {
		if grid[y][i+x+1] >= height {
			right = false
			break
		}
	}
	if right {
		return true
	}
	// check up
	for _, row := range grid[:y] {
		if row[x] >= height {
			up = false
			break
		}
	}
	if up {
		return true
	}
	// check down
	for _, row := range grid[y+1:] {
		if row[x] >= height {
			down = false
			break
		}
	}
	return down
}
