package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMask_ApplyToValue(t *testing.T) {
	for _, test := range []struct {
		Mask Mask
		Num  int
		Exp  int
	}{
		{Mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"), 11, 73},
		{Mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"), 101, 101},
		{Mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"), 0, 64},
		{Mask("011011X11X11100101XX0XX0100100000X0X"), 171994, 29453846784},
	} {
		got, err := test.Mask.ApplyToValue(test.Num)
		if err != nil {
			t.Fatal(err)
		}
		if got != test.Exp {
			t.Errorf("apply %d to %v: expected %d, got %d", test.Num, test.Mask, test.Exp, got)
		}
	}
}

func TestMask_ApplyToAddress(t *testing.T) {
	for _, test := range []struct {
		Mask Mask
		Num  int
		Exp  []int
	}{
		{Mask("000000000000000000000000000000X1001X"), 42, []int{26, 27, 58, 59}},
	} {
		got, err := test.Mask.ApplyToAddress(test.Num)
		if err != nil {
			t.Fatal(err)
		}
		sort.Ints(got)
		sort.Ints(test.Exp)
		if !reflect.DeepEqual(got, test.Exp) {
			t.Errorf("apply %d to %v: expected %d, got %d", test.Num, test.Mask, test.Exp, got)
		}
	}
}

func TestProgram_Sum(t *testing.T) {
	program := Program{
		7: 101,
		8: 64,
	}
	exp := 165
	if got := program.Sum(); got != exp {
		t.Errorf("program.Sum() = %d but got %d", exp, got)
	}
}
