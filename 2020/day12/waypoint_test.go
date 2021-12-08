package main

import (
	"reflect"
	"testing"
)

func TestWaypoint_turnRight(t *testing.T) {
	for _, tt := range []struct {
		wp   Waypoint
		want Waypoint
	}{
		{Waypoint{1, 5}, Waypoint{-5, 1}},
	} {
		t.Logf("turnRight() on %v", tt.wp)
		tt.wp.turnRight()
		if !reflect.DeepEqual(tt.wp, tt.want) {
			t.Errorf("turnRight() = %v, want %v", tt.wp, tt.want)
		}
	}
}
