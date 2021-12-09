package main

import (
	"fmt"
	"strings"
	"sync"

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
	// fmt.Println(part2(input))
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

func countLists(lists [][]int) int {
	res := 0
	for _, l := range lists {
		res += len(l)
	}
	return res
}

func reproduce(fish []int, days int) int {
	var schools [][]int
	for _, f := range fish {
		schools = append(schools, []int{f})
	}
	for d := 0; d < days; d++ {
		fmt.Println(d, len(schools), len(schools[0]), countLists(schools))
		var wg sync.WaitGroup
		for si := range schools {
			si := si
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := range schools[si] {
					schools[si][i]--
					if schools[si][i] == -1 {
						schools[si] = append(schools[si], newAge)
						schools[si][i] = age
					}
				}
			}()
		}
		wg.Wait()
	}
	return countLists(schools)
}
