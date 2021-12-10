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
	fmt.Println(part2(input))
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
	var e Entry
	for _, s := range strings.Fields(split[0]) {
		e.Patterns = append(e.Patterns, sortString(s))
	}
	for _, s := range strings.Fields(split[1]) {
		e.Output = append(e.Output, sortString(s))
	}
	return e
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

func part2(entries []Entry) int {
	var res int
	for _, e := range entries {
		res += e.ResolveOutput()
	}
	return res
}

func (e Entry) ResolveOutput() int {
	var res int
	if l := len(e.Output); l != 4 {
		panic(fmt.Sprintf("output should be 4 words long but is %d: %v", l, e.Output))
	}
	mappings := e.FindPatterns()
	for i, o := range e.Output {
		d := mappings[o]
		switch i {
		case 0:
			res += d * 1000
		case 1:
			res += d * 100
		case 2:
			res += d * 10
		case 3:
			res += d
		}
	}
	return res
}

func (e Entry) FindPatterns() map[string]int {
	var zero, one, two, three, four, five, six, seven, eight, nine string
	var zeroSixNine, twoThreeFive []string
	for _, val := range e.Patterns {
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

	return map[string]int{
		zero:  0,
		one:   1,
		two:   2,
		three: 3,
		four:  4,
		five:  5,
		six:   6,
		seven: 7,
		eight: 8,
		nine:  9,
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

func sortStrings(ix []string) []string {
	var res []string
	for i := range ix {
		res = append(res, sortString(ix[i]))
	}
	return res
}

func sortString(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		panic("sortString with empty string not allowed")
	}
	var sx []string
	for _, r := range input {
		sx = append(sx, string(r))
	}
	sort.Strings(sx)
	return strings.Join(sx, "")
}
