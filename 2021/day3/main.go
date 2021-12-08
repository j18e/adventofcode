package main

import (
	"fmt"

	"github.com/j18e/adventofcode/pkg/converting"
	"github.com/j18e/adventofcode/pkg/inputting"
)

func main() {
	input := inputting.GetInputStrings("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	chars := len(input[0])

	var gammaStr, epsilonStr string
	for i := 0; i < chars; i++ {
		zeroes, ones := totals(input, i)
		if zeroes == ones {
			panic("zeroes and ones are the same!")
		}
		if zeroes > ones {
			gammaStr += "0"
			epsilonStr += "1"
		} else {
			gammaStr += "1"
			epsilonStr += "0"
		}
	}
	gamma := converting.Btoi(gammaStr)
	epsilon := converting.Btoi(epsilonStr)
	return int(gamma) * int(epsilon)
}

func part2(input []string) int {
	chars := len(input[0])

	// calculate o2
	o2Lines := append([]string{}, input...)
	for i := 0; i < chars; i++ {
		zeroes, ones := totals(o2Lines, i)
		if ones >= zeroes {
			o2Lines = filterByBit(o2Lines, i, '1')
		} else {
			o2Lines = filterByBit(o2Lines, i, '0')
		}
		if len(o2Lines) == 1 {
			break
		}
	}
	if l := len(o2Lines); l != 1 {
		panic(fmt.Sprintf("got %d o2 results, expected 1", l))
	}
	o2 := converting.Btoi(o2Lines[0])

	// calculate co2
	co2Lines := append([]string{}, input...)
	for i := 0; i < chars; i++ {
		zeroes, ones := totals(co2Lines, i)
		if zeroes <= ones {
			co2Lines = filterByBit(co2Lines, i, '0')
		} else {
			co2Lines = filterByBit(co2Lines, i, '1')
		}
		if len(co2Lines) == 1 {
			break
		}
	}
	if l := len(co2Lines); l != 1 {
		panic(fmt.Sprintf("got %d co2 results, expected 1", l))
	}
	co2 := converting.Btoi(co2Lines[0])

	return o2 * co2
}

func filterByBit(input []string, pos int, bit rune) []string {
	var res []string
	for _, ln := range input {
		if rune(ln[pos]) == bit {
			res = append(res, ln)
		}
	}
	return res
}

func totals(input []string, char int) (zeroes, ones int) {
	for _, ln := range input {
		switch ln[char] {
		case '0':
			zeroes++
		case '1':
			ones++
		default:
			panic("expected '0' or '1': " + ln)
		}
	}
	return zeroes, ones
}
