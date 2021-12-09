package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/j18e/adventofcode/pkg/inputting"
)

type Digit string

const (
	Zero  Digit = "abcefg"
	One         = "cf"
	Two         = "acdeg"
	Three       = "acdfg"
	Four        = "bcdf"
	Five        = "abdfg"
	Six         = "abdefg"
	Seven       = "acf"
	Eight       = "abcdefg"
	Nine        = "abcdfg"
)

func main() {
	input := ParseEntries(inputting.GetInputStrings("input.txt"))
	fmt.Println(part1(input))
	// fmt.Println(part2(input))
}

func ParseEntries(input []string) []Entry {
	var entries []Entry
	for _, ln := range input {
		split := strings.Split(ln, "|")
		entries = append(entries, Entry{
			Patterns: strings.Split(split[0], " "),
			Output:   strings.Split(split[1], " "),
		})
	}
	return entries
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
	return 0
}

type Mappings struct {
	T, TL, TR, M, BL, BR, B string
}

func (m Mappings) One() string {
	return m.TR + m.BR
}

func (m Mappings) Two() string {
	return m.T + m.TR + m.M + m.BL + m.B
}

func (m Mappings) Three() string {
	return m.T + m.TR + m.M + m.BR + m.B
}

func (m Mappings) Four() string {
	return m.TL + m.TR + m.M + m.BR
}

func (m Mappings) Five() string {
	return m.T + m.TL + m.M + m.BR + m.B
}

func (m Mappings) Six() string {
	return m.T + m.TL + m.M + m.BL + m.BR + m.B
}

func (m Mappings) Seven() string {
	return m.T + m.TR + m.BR
}

func (m Mappings) Eight() string {
	return m.T + m.TL + m.TR + m.M + m.BL + m.BR + m.B
}

func (m Mappings) Nine() string {
	return m.T + m.TL + m.TR + m.M + m.BR + m.B
}

// func deriveMappings(input []string) Mappings {
// 	sortStrings(input)
// 	lengths := make(map[int][]string)
// 	for _, i := range input {
// 		lengths[len(i)] = append(lengths[len(i)], i)
// 	}

// 	return mappings
// }

func allCombinations() []string {
	const letters = "abcdefg"
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
