package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
	"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
	"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
	"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
	"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
	"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
	"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
	"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
	"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
	"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
}

var singleEntry = Entry{
	sortStrings([]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}),
	sortStrings([]string{"cdfeb", "fcadb", "cdfeb", "cdbaf"}),
}

func Test_part1(t *testing.T) {
	exp := 26
	got := part1(ParseEntries(input))
	assert.Equal(t, exp, got)
}

func Test_part2(t *testing.T) {
	exp := 61229
	got := part2(ParseEntries(input))
	assert.Equal(t, exp, got)
}

func Test_sortString(t *testing.T) {
	assert.Equal(t, "abcdefg", sortString("cfbegad"))
}

func Test_commonLetters(t *testing.T) {
	for _, tt := range []struct {
		d1, d2 string
		exp    int
	}{
		{"abc", "", 0},
		{"abc", "cab", 3},
		{"", "abc", 0},
	} {
		assert.Equal(t, tt.exp, commonLetters(tt.d1, tt.d2))
	}
}

func TestEntry_ResolveOutput(t *testing.T) {
	assert := assert.New(t)
	{
		exp := 5353
		res := singleEntry.ResolveOutput()
		assert.Equal(exp, res)
	}

	exp := []int{8394, 9781, 1197, 9361, 4873, 8418, 4548, 1625, 8717, 4315}
	assert.Equal(len(exp), len(input))
	for i := range exp {
		e := ParseEntry(input[i])
		assert.Equal(exp[i], e.ResolveOutput())
	}
}

func TestEntry_FindPatterns(t *testing.T) {
	exp := map[string]int{
		"abcdeg":  0,
		"ab":      1,
		"acdfg":   2,
		"abcdf":   3,
		"abef":    4,
		"bcdef":   5,
		"bcdefg":  6,
		"abd":     7,
		"abcdefg": 8,
		"abcdef":  9,
	}
	for s, n := range exp {
		delete(exp, s)
		exp[sortString(s)] = n
	}
	got := singleEntry.FindPatterns()
	assert.Equal(t, exp, got)
}
