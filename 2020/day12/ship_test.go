package main

import (
	"testing"
)

func TestShip_DistanceFromStart(t *testing.T) {
	for _, tt := range []struct {
		ix   []Instruction
		want int
	}{
		{[]Instruction{}, 0},
		{[]Instruction{{N, 10}, {R, 90}, {F, 10}}, 0},
	} {
		s := &Ship{Heading: E}
		s.FollowCourse(tt.ix)
		if got := s.DistanceFromStart(); got != tt.want {
			t.Errorf("Ship.DistanceFromStart() = %d, want %d", got, tt.want)
		}
	}
}

func TestShip_MoveTowardsWaypoint(t *testing.T) {
	for _, tt := range []struct {
		wp    *Waypoint
		times int
		want  int
	}{
		{&Waypoint{1, 5}, 5, 30},
		{&Waypoint{10, 10}, 10, 200},
		{&Waypoint{-10, 1}, 2, 22},
	} {
		s := &Ship{}
		s.MoveTowardsWaypoint(tt.wp, tt.times)
		if got := s.DistanceFromStart(); got != tt.want {
			t.Errorf("Ship.DistanceFromStart() = %d, want %d", got, tt.want)
		}
	}
}

func TestShip_FollowCourseWithWaypoint(t *testing.T) {
	for _, tt := range []struct {
		wp   *Waypoint
		ix   []Instruction
		want int
	}{
		// ship shouldn't move at all if there is no F instruction
		{&Waypoint{1, 5}, []Instruction{{N, 10}}, 0},

		{&Waypoint{1, 5}, []Instruction{{F, 5}}, 30},
		{&Waypoint{1, 5}, []Instruction{{N, 9}, {F, 5}}, 75},

		{&Waypoint{1, 5}, []Instruction{{N, 9}, {R, 90}, {F, 5}}, 75},
	} {
		s := &Ship{}
		s.FollowCourseWithWaypoint(tt.wp, tt.ix)
		if got := s.DistanceFromStart(); got != tt.want {
			t.Errorf("Ship.DistanceFromStart() = %d, want %d", got, tt.want)
		}
	}
}
