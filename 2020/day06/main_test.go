package main

import "testing"

func TestPart1(t *testing.T) {
	groups := []string{`abc`, "a\nb\nc", "ab\nac", "a\na\na\na", "b", }
	exp := 11
	got := part1(groups)
	if got != exp {
		t.Errorf("sum %v: expected %d, got %d", groups, exp, got)
	}
}

func TestGroup_SumAllYes(t *testing.T) {
	g := Group{
		Members: 3,
		Letters: map[rune]int{ 'a': 3, 'b': 2, 'c': 3, 'd': 4, },
	}

	exp := 2
	if got := g.SumAllYes(); got != exp {
		t.Errorf("sum %v: expected %d, got %d", g, exp, got)
	}
}

