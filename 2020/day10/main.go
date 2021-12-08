package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var nums []int
	for _, ln := range strings.Split(string(bs), "\n") {
		if ln == "" {
			continue
		}
		num, err := strconv.Atoi(ln)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	pf := PathFinder{
		Counts:   make(map[int]int),
		Joltages: assembleJoltages(nums),
	}

	ones, threes := pf.SumDifferences()
	fmt.Println("part 1:", ones*threes)
	fmt.Println("part 2:", pf.ValidPaths(0))
}

func assembleJoltages(nums []int) map[int]bool {
	res := map[int]bool{0: true}
	var highest int
	for _, num := range nums {
		if num > highest {
			highest = num
		}
		res[num] = true
	}
	res[highest+3] = true
	return res
}

type PathFinder struct {
	Counts   map[int]int
	Joltages map[int]bool
}

func (p PathFinder) ValidPaths(start int) int {
	if cnt, ok := p.Counts[start]; ok {
		return cnt
	}
	paths := 0
	for i := start + 1; i <= start+3; i++ {
		if p.Joltages[i] {
			paths += p.ValidPaths(i)
		}
	}
	if paths == 0 {
		paths++
	}
	p.Counts[start] = paths
	return paths
}

func (p PathFinder) SumDifferences() (ones, threes int) {
	for num := range p.Joltages {
		if p.Joltages[num+1] {
			ones++
			continue
		}
		if p.Joltages[num+3] {
			threes++
			continue
		}
	}
	return ones, threes
}
