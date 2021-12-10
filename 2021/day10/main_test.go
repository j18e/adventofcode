package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

func Test_part2Score(t *testing.T) {
	for _, tt := range []struct {
		cs  string
		exp int
	}{
		{"}}]])})]", 288957},
		{")}>]})", 5566},
		{"}}>}>))))", 1480781},
		{"]]}}]}]}>", 995444},
		{"])}>", 294},
	} {
		assert.Equal(t, tt.exp, part2Score(tt.cs))
	}
}

func Test_part1Score(t *testing.T) {
	inp := []rune{'}', ')', ']', ')', '>'}
	assert.Equal(t, 26397, part1Score(inp))
}

func Test_completeLine(t *testing.T) {
	for i, exp := range []Line{
		{false, 0, "}}]])})]"},
		{false, 0, ")}>]})"},
		{true, '}', ""},
		{false, 0, "}}>}>))))"},
		{true, ')', ""},
		{true, ']', ""},
		{false, 0, "]]}}]}]}>"},
		{true, ')', ""},
		{true, '>', ""},
		{false, 0, "])}>"},
	} {
		assert.Equal(t, exp, completeLine(input[i]))
	}
}
