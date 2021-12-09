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
	"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`",
}

func Test_part1(t *testing.T) {
	exp := 26
	got := part1(ParseEntries(input))
	assert.Equal(t, exp, got)
}

func Test_sortString(t *testing.T) {
	exp := "abc"
	got := sortString("bca")
	assert.Equal(t, exp, got)
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

func Test_findPatterns(t *testing.T) {
	input := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	sortStrings(input)
	exp := []string{"cagedb", "ab", "gcdfa", "fbcad", "eafb", "cdfbe", "cdfgeb", "dab", "acedgfb", "cefabd"}
	sortStrings(exp)
	got := findPatterns(input)
	assert.Equal(t, exp, got)
}
