package main

import (
	"fmt"
	"log"

	"github.com/j18e/adventofcode/pkg/input"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := input.GetInput("input.txt")
	part1(input)
	part2(input)
	return nil
}

func part1(input []int) {
	res := 0
	// start at idx 1 because idx 0 has no previous measurement
	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			res++
		}
	}
	fmt.Println(res)
}

func part2(input []int) {
	res := 0
	// we need to stop iterating before the end to avoid panic
	for i := 1; i < len(input)-2; i++ {
		window := sum(input[i], input[i+1], input[i+2])
		prevWindow := sum(input[i-1], input[i], input[i+1])
		if window > prevWindow {
			res++
		}
	}
	fmt.Println(res)
}

func sum(ix ...int) int {
	res := 0
	for _, i := range ix {
		res += i
	}
	return res
}
