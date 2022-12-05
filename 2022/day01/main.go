package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bs, err := os.ReadFile("input.txt")
	if err != nil {
		return err
	}
	elves := [][]int{{}}
	for _, ln := range strings.Split(string(bs), "\n") {
		if ln == "" {
			elves = append(elves, []int{})
			continue
		}
		num, err := strconv.Atoi(ln)
		if err != nil {
			return err
		}
		elves[len(elves)-1] = append(elves[len(elves)-1], num)
	}
	var totals []int
	for _, elf := range elves {
		totals = append(totals, sum(elf))
	}
	sort.Slice(totals, func(i, j int) bool { return totals[i] > totals[j] })
	fmt.Println(totals[0])
	fmt.Println(totals[0] + totals[1] + totals[2])
	return nil
}

func sum(ix []int) int {
	sum := 0
	for _, i := range ix {
		sum += i
	}
	return sum
}
