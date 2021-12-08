package main

import (
	"testing"
)

var testGrid *Grid

func fillTestGrid() {
	o, e, f := SpotOccupied, SpotEmpty, SpotFloor
	testGrid = &Grid{
		3, 3,
		map[int]Spot{
			0: o, 1: o, 2: o,
			3: o, 4: e, 5: o,
			6: f, 7: o, 8: o,
		},
	}
}

func TestGrid_seatOccupied(t *testing.T) {
	fillTestGrid()
	for _, tt := range []struct {
		inp  coordinate
		want bool
	}{
		{coordinate{0, 0}, true},
		{coordinate{1, 1}, false},
		{coordinate{-1, -1}, false},
	} {
		if got := testGrid.seatOccupied(tt.inp); got != tt.want {
			t.Errorf("Grid.seatOccupied() = %v, want %v", got, tt.want)
		}
	}
}

func TestGrid_CountAdjacentOccupied(t *testing.T) {
	fillTestGrid()
	for _, tt := range []struct {
		inp, want int
	}{
		{4, 7},
		{0, 2},
		{8, 2},
		{7, 3},
	} {
		if got := testGrid.CountAdjacentOccupied(tt.inp); got != tt.want {
			t.Errorf("Grid.CountAdjacentOccupied(%d) = %d, want %d", tt.inp, got, tt.want)
		}
	}
}

func TestGrid_SeatFromID(t *testing.T) {
	fillTestGrid()
	wantRow := 1
	wantSeat := 1
	row, seat := testGrid.SeatFromID(4)
	if row != wantRow || seat != wantSeat {
		t.Errorf("Grid.SeatFromID() = %d %d, want %d %d", row, seat, wantRow, wantSeat)
	}
}

func TestGrid_CountOccupiedAllDirections(t *testing.T) {
	fillTestGrid()
	tests := []struct {
		id, want int
	}{
		{0, 2},
		{3, 3},
	}
	for _, tt := range tests {
		if got := testGrid.CountOccupiedAllDirections(tt.id); got != tt.want {
			t.Errorf("Grid.CountOccupiedAllDirections() = %v, want %v", got, tt.want)
		}
	}
}
