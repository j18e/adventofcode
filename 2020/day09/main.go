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

	var lines []int
	for _, ln := range strings.Split(string(bs), "\n") {
		if ln == "" {
			continue
		}
		i, err := strconv.Atoi(ln)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, i)
	}

	num := firstNotContainingSum(lines, 25)
	fmt.Println("part 1:", num)
	fmt.Println("part 2:", sumTo(num, lines))
}

func firstNotContainingSum(lines []int, preamble int) int {
	for i := preamble; i < len(lines); i++ {
		if !containsSum(lines[i], lines[i-preamble:i]) {
			return lines[i]
		}
	}
	return -1
}

func containsSum(num int, ix []int) bool {
	if len(ix) < 2 {
		return false
	}
	for i := 1; i < len(ix); i++ {
		if ix[0]+ix[i] == num {
			return true
		}
	}
	return containsSum(num, ix[1:])
}

func sumTo(num int, ix []int) int {
	if len(ix) < 2 {
		return -1
	}
	tot := ix[0]
	i := 1
	for ; i < len(ix); i++ {
		tot += ix[i]
		if tot >= num {
			break
		}
	}
	if tot == num {
		return sumSmallestAndLargest(ix[:i+1])
	}
	return sumTo(num, ix[1:])
}

func sumSmallestAndLargest(ix []int) int {
	small, large := ix[0], ix[0]
	for _, n := range ix {
		if n < small {
			small = n
		}
		if n > large {
			large = n
		}
	}
	return small + large
}
