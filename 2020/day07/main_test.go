package main

import "testing"

func TestPart2(t *testing.T) {
	input := BagList{
		"gold":   map[string]int{"red": 2},
		"red":    map[string]int{"orange": 2},
		"orange": map[string]int{"yellow": 2},
		"yellow": map[string]int{"green": 2},
		"green":  map[string]int{"blue": 2},
		"blue":   map[string]int{"violet": 2},
		"violet": nil,
	}

	exp := 127 // including the gold bag itself
	if got := input.CountBags("gold"); got != exp {
		t.Errorf("got %d, expected %d", got, exp)
	}
}
