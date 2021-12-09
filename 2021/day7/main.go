package main

import (
	"fmt"
	"strings"

	"github.com/j18e/adventofcode/pkg/converting"
	"github.com/j18e/adventofcode/pkg/inputting"
)

func main() {
	crabs := parseInput(inputting.GetInput("input.txt"))
	fmt.Println(part1(crabs))
	fmt.Println(part2(crabs))
}

func parseInput(input string) []int {
	var res []int
	for _, num := range strings.Split(input, ",") {
		res = append(res, converting.Atoi(num))
	}
	return res
}

func part1(crabs []int) int {
	return bestDestination(crabs, distance)
}

func part2(crabs []int) int {
	return bestDestination(crabs, requiredFuel)
}

func bestDestination(crabs []int, fuelCalc func(int, int) int) int {
	min, max := minMax(crabs)
	fuelReq := make(map[int]int)
	for i := min; i <= max; i++ {
		for _, pos := range crabs {
			fuelReq[i] += fuelCalc(pos, i)
		}
	}
	leastFuel := fuelReq[min]
	for _, fuel := range fuelReq {
		if fuel < leastFuel {
			leastFuel = fuel
		}
	}
	return leastFuel
}

func requiredFuel(pos1, pos2 int) int {
	res := 0
	for i := 1; i <= distance(pos1, pos2); i++ {
		res += i
	}
	return res
}

func distance(pos1, pos2 int) int {
	if pos1 == pos2 {
		return 0
	}
	if pos1 > pos2 {
		return pos1 - pos2
	}
	return pos2 - pos1
}

func minMax(ix []int) (int, int) {
	min, max := ix[0], ix[0]
	for _, i := range ix {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return min, max
}
