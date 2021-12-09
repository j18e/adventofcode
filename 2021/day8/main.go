package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/j18e/adventofcode/pkg/inputting"
)

func main() {
	input := ParseEntries(inputting.GetInputStrings("input.txt"))
	fmt.Println(part1(input))
	// fmt.Println(part2(input))
}

func ParseEntries(input []string) []Entry {
	var entries []Entry
	for _, ln := range input {
		entries = append(entries, ParseEntry(ln))
	}
	return entries
}

func ParseEntry(input string) Entry {
	split := strings.Split(input, "|")
	return Entry{
		Patterns: strings.Split(split[0], " "),
		Output:   strings.Split(split[1], " "),
	}
}

type Entry struct {
	Patterns []string
	Output   []string
}

func part1(entries []Entry) int {
	var res int
	for _, e := range entries {
		for _, digit := range e.Output {
			switch len(digit) {
			case 2, 3, 4, 7:
				res++
			}
		}
	}
	return res
}

var lengths = map[int][]int{
	2: {1},
	3: {7},
	4: {4},
	5: {2, 3, 5},
	6: {0, 6, 9},
	7: {8},
}

func part2(entries []Entry) int {
	return 0
}

func findPatterns(patterns []string) []string {
	sortStrings(patterns)
	var zero, one, two, three, four, five, six, seven, eight, nine string
	var zeroSixNine, twoThreeFive []string
	for _, val := range patterns {
		switch len(val) {
		case 2:
			one = val
		case 3:
			seven = val
		case 4:
			four = val
		case 5:
			twoThreeFive = append(twoThreeFive, val)
		case 6:
			zeroSixNine = append(zeroSixNine, val)
		case 7:
			eight = val
		}
	}

	for _, val := range zeroSixNine {
		if commonLetters(val, four) == 4 {
			nine = val
		} else {
			if commonLetters(val, one) == 2 {
				zero = val
			} else {
				six = val
			}
		}
	}

	for _, val := range twoThreeFive {
		if commonLetters(val, one) == 2 {
			three = val
		} else {
			if commonLetters(val, four) == 2 {
				two = val
			} else {
				five = val
			}
		}
	}

	return []string{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}
}

func commonLetters(d1, d2 string) int {
	var res int
	for _, l := range d1 {
		for _, ll := range d2 {
			if l == ll {
				res++
			}
		}
	}
	return res
}

func sortStrings(ix []string) {
	for i := range ix {
		ix[i] = sortString(ix[i])
	}
}

func sortString(input string) string {
	split := strings.Split(input, "")
	sort.Strings(split)
	return strings.Join(split, "")
}
