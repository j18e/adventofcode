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
	var vals []int
	for _, ln := range input {
		split := strings.Split(ln, " ")
		switch len(split) {
		case 1:
			vals = append(vals, 0)
		case 2:
			num, err := strconv.Atoi(split[1])
			if err != nil {
				return err
			}
			vals = append(vals, num)
		default:
			panic(ln)
		}
	}

	steps := FindSteps(vals, 1)
	part1(steps)
	steps = steps[1:]
	part2(steps, 6)
	return nil
}

func part2(steps []int, lines int) {
	var res string
	for i := 0; i < lines; i++ {
		res += addPixels(steps[i*40 : i*40+40])
	}
	fmt.Println(strings.TrimSpace(res))
}

func addPixels(steps []int) string {
	res := ""
	for i, x := range steps {
		res = addPixel(res, i, x)
	}
	return res + "\n"
}

func addPixel(res string, i, x int) string {
	switch i {
	case x - 1, x, x + 1:
		res += "#"
	default:
		res += "."
	}
	return res
}

func part1(steps []int) {
	total := 0
	for i := 0; i < 6; i++ {
		cycle := 20 + i*40
		val := steps[cycle]
		total += cycle * val
	}
	fmt.Println(total)
}

func FindSteps(vals []int, x int) []int {
	res := []int{x, x}
	i := 0
	for _, val := range vals {
		res = append(res, x)
		i++
		if val == 0 {
			continue
		}
		x += val
		res = append(res, x)
	}
	return res
}

var input = []string{
	`addx 15`,
	`addx -11`,
	`addx 6`,
	`addx -3`,
	`addx 5`,
	`addx -1`,
	`addx -8`,
	`addx 13`,
	`addx 4`,
	`noop`,
	`addx -1`,
	`addx 5`,
	`addx -1`,
	`addx 5`,
	`addx -1`,
	`addx 5`,
	`addx -1`,
	`addx 5`,
	`addx -1`,
	`addx -35`,
	`addx 1`,
	`addx 24`,
	`addx -19`,
	`addx 1`,
	`addx 16`,
	`addx -11`,
	`noop`,
	`noop`,
	`addx 21`,
	`addx -15`,
	`noop`,
	`noop`,
	`addx -3`,
	`addx 9`,
	`addx 1`,
	`addx -3`,
	`addx 8`,
	`addx 1`,
	`addx 5`,
	`noop`,
	`noop`,
	`noop`,
	`noop`,
	`noop`,
	`addx -36`,
	`noop`,
	`addx 1`,
	`addx 7`,
	`noop`,
	`noop`,
	`noop`,
	`addx 2`,
	`addx 6`,
	`noop`,
	`noop`,
	`noop`,
	`noop`,
	`noop`,
	`addx 1`,
	`noop`,
	`noop`,
	`addx 7`,
	`addx 1`,
	`noop`,
	`addx -13`,
	`addx 13`,
	`addx 7`,
	`noop`,
	`addx 1`,
	`addx -33`,
	`noop`,
	`noop`,
	`noop`,
	`addx 2`,
	`noop`,
	`noop`,
	`noop`,
	`addx 8`,
	`noop`,
	`addx -1`,
	`addx 2`,
	`addx 1`,
	`noop`,
	`addx 17`,
	`addx -9`,
	`addx 1`,
	`addx 1`,
	`addx -3`,
	`addx 11`,
	`noop`,
	`noop`,
	`addx 1`,
	`noop`,
	`addx 1`,
	`noop`,
	`noop`,
	`addx -13`,
	`addx -19`,
	`addx 1`,
	`addx 3`,
	`addx 26`,
	`addx -30`,
	`addx 12`,
	`addx -1`,
	`addx 3`,
	`addx 1`,
	`noop`,
	`noop`,
	`noop`,
	`addx -9`,
	`addx 18`,
	`addx 1`,
	`addx 2`,
	`noop`,
	`noop`,
	`addx 9`,
	`noop`,
	`noop`,
	`noop`,
	`addx -1`,
	`addx 2`,
	`addx -37`,
	`addx 1`,
	`addx 3`,
	`noop`,
	`addx 15`,
	`addx -21`,
	`addx 22`,
	`addx -6`,
	`addx 1`,
	`noop`,
	`addx 2`,
	`addx 1`,
	`noop`,
	`addx -10`,
	`noop`,
	`noop`,
	`addx 20`,
	`addx 1`,
	`addx 2`,
	`addx 2`,
	`addx -6`,
	`addx -11`,
	`noop`,
	`noop`,
	`noop`,
}
