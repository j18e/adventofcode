package main

import (
	"fmt"
	"sort"

	"github.com/j18e/adventofcode/pkg/inputting"
)

func main() {
	input := inputting.GetInputStrings("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	var offenders []rune
	for _, ln := range input {
		line := completeLine(ln)
		if line.Corrupt {
			offenders = append(offenders, line.OffendingChar)
		}
	}
	return part1Score(offenders)
}

func part2(input []string) int {
	var scores []int
	for _, ln := range input {
		line := completeLine(ln)
		if !line.Corrupt {
			scores = append(scores, part2Score(line.Closers))
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func part2Score(cs string) int {
	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	score := 0
	for _, r := range cs {
		score *= 5
		score += points[r]
	}
	return score
}

func part1Score(rx []rune) int {
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	var res int
	for _, r := range rx {
		res += points[r]
	}
	return res
}

type Line struct {
	Corrupt       bool
	OffendingChar rune
	Closers       string
}

func completeLine(line string) Line {
	var closers string
	for _, r := range line {
		if isOpener(r) {
			closers = opposite(r) + closers
			continue
		}
		if r != rune(closers[0]) {
			return Line{true, r, ""}
		}
		closers = closers[1:]
	}
	return Line{Closers: closers}
}

func isOpener(r rune) bool {
	switch r {
	case '(', '[', '{', '<':
		return true
	}
	return false
}

func opposite(r rune) string {
	switch r {
	case '(':
		return ")"
	case '[':
		return "]"
	case '{':
		return "}"
	case '<':
		return ">"
	}
	panic(fmt.Sprintf("unknown rune %s", string(r)))
}
