package main

import (
	"fmt"
	"strings"

	"github.com/j18e/adventofcode/pkg/converting"
	"github.com/j18e/adventofcode/pkg/inputting"
)

const (
	age    = 6
	newAge = 8
)

func main() {
	input := parseInput(inputting.GetInput("input.txt"))
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(fish []int) int {
	return reproduce(fish, 80)
}

func part2(fish []int) int {
	return reproduce(fish, 256)
}

func parseInput(input string) []int {
	var res []int
	for _, num := range strings.Split(input, ",") {
		res = append(res, converting.Atoi(num))
	}
	return res
}

func countFish(fish map[int]int) int {
	res := 0
	for _, cnt := range fish {
		res += cnt
	}
	return res
}

func reproduce(startingFish []int, days int) int {
	fish := make(map[int]int, 9)
	for _, f := range startingFish {
		fish[f]++
	}

	for d := 0; d < days; d++ {
		fish = map[int]int{
			0: fish[1],
			1: fish[2],
			2: fish[3],
			3: fish[4],
			4: fish[5],
			5: fish[6],
			6: fish[7] + fish[0],
			7: fish[8],
			8: fish[0],
		}
	}
	return countFish(fish)
}
