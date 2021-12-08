package main

import (
	"testing"
)

func TestFirstNotContainingSum(t *testing.T) {
	for _, test := range []struct {
		pre, exp int
		inp      []int
	}{
		{5, 127, []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}},
		{5, 99, []int{1, 2, 3, 4, 5, 6, 99}},
	} {
		if got := firstNotContainingSum(test.inp, test.pre); got != test.exp {
			t.Errorf("expected %d, got %d", test.exp, got)
		}
	}

}

func TestContainsSum(t *testing.T) {
	for _, test := range []struct {
		num int
		exp bool
		inp []int
	}{
		{5, true, []int{2, 3}},
		{5, false, []int{2, 2}},
	} {
		if got := containsSum(test.num, test.inp); got != test.exp {
			t.Errorf("containsSum(%d, %v): expected %t, got %t", test.num, test.inp, test.exp, got)
		}
	}

}

func TestSumTo(t *testing.T) {
	for _, test := range []struct {
		num, exp int
		inp      []int
	}{
		{6, 4, []int{1, 2, 3}},
		{5, 5, []int{1, 2, 3, 4, 5, 6}},
		{9, 6, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	} {
		if got := sumTo(test.num, test.inp); got != test.exp {
			t.Errorf("sumTo(%d, %v): expected %d, got %d", test.num, test.inp, test.exp, got)
		}
	}

}
