package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/j18e/adventofcode/2022/common"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := common.ReadInput("input.txt")
	var moves []Move
	for _, ln := range input {
		split := strings.Split(ln, " ")
		cnt, err := strconv.Atoi(split[1])
		if err != nil {
			return err
		}
		moves = append(moves, Move{split[0], cnt})
	}

	grid1 := &Grid{
		Knots:       []Coord{{}, {}},
		TailVisited: make(map[string]bool),
	}
	grid1.MarkTailVisited()

	grid2 := &Grid{TailVisited: make(map[string]bool)}
	for i := 0; i < 10; i++ {
		grid2.Knots = append(grid2.Knots, Coord{})
	}
	grid2.MarkTailVisited()

	for _, m := range moves {
		for i := 0; i < m.Count; i++ {
			switch m.Direction {
			case "L":
				grid1.Knots[0].x--
				grid2.Knots[0].x--
			case "R":
				grid1.Knots[0].x++
				grid2.Knots[0].x++
			case "U":
				grid1.Knots[0].y--
				grid2.Knots[0].y--
			case "D":
				grid1.Knots[0].y++
				grid2.Knots[0].y++
			default:
				return fmt.Errorf("unrecognized direction %s", m.Direction)
			}
			grid1.Follow(1)
			for i := range grid2.Knots {
				if i == 0 {
					continue
				}
				grid2.Follow(i)
			}
		}
	}
	fmt.Println("part 1:", len(grid1.TailVisited))
	fmt.Println("part 2:", len(grid2.TailVisited))

	return nil
}

type Move struct {
	Direction string
	Count     int
}

func Diff(head, tail int) (diff int, minus int) {
	d := head - tail
	if d < 0 {
		return d * -1, -1
	}
	return d, 1
}

type Grid struct {
	Knots       []Coord
	TailVisited map[string]bool
	TailLog     []string
}

func (g *Grid) String() string {
	res := ""
	for _, k := range g.Knots {
		res += fmt.Sprintf("%d,%d..", k.x, k.y)
	}
	return strings.TrimRight(res, ".")
}

func (g *Grid) MarkTailVisited() {
	tail := g.Knots[len(g.Knots)-1].String()
	if !g.TailVisited[tail] {
		g.TailVisited[tail] = true
		g.TailLog = append(g.TailLog, tail)
	}
}

func (g *Grid) Follow(i int) bool {
	xDiff, xMinus := Diff(g.Knots[i-1].x, g.Knots[i].x)
	yDiff, yMinus := Diff(g.Knots[i-1].y, g.Knots[i].y)
	switch {
	case xDiff <= 1 && yDiff <= 1: // no adjustment required
		return false
	case xDiff == 0: // only change y
		g.Knots[i].y += (yDiff - 1) * yMinus
	case yDiff == 0: // only change x
		g.Knots[i].x += (xDiff - 1) * xMinus
	case xDiff != yDiff: // uneven diagonal
		g.Knots[i].x += 1 * xMinus
		g.Knots[i].y += 1 * yMinus
	default: // even diagonal
		g.Knots[i].x += (xDiff - 1) * xMinus
		g.Knots[i].y += (yDiff - 1) * yMinus
	}
	g.MarkTailVisited()
	return true
}

type Coord struct {
	x, y int
}

func (c Coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}
