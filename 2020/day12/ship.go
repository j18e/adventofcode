package main

import (
	"math"
)

type Ship struct {
	Heading Direction
	NS, EW  int
}

func (s *Ship) MoveTowardsWaypoint(wp *Waypoint, times int) {
	for i := 0; i < times; i++ {
		s.NS += wp.NS
		s.EW += wp.EW
	}
}

func (s *Ship) FollowCourseWithWaypoint(wp *Waypoint, ix []Instruction) {
	for _, ins := range ix {
		switch ins.D {
		case L, R:
			wp.Turn(ins.D, ins.N/90)
		case N, S, E, W:
			wp.Move(ins.D, ins.N)
		case F:
			s.MoveTowardsWaypoint(wp, ins.N)
		}
	}
}

func (s *Ship) FollowCourse(ix []Instruction) {
	for _, ins := range ix {
		switch ins.D {
		case N, S, E, W:
			s.Move(ins.D, ins.N)
		case F:
			s.Move(s.Heading, ins.N)
		case L:
			s.Heading = turnLeft(s.Heading, ins.N/90)
		case R:
			s.Heading = turnRight(s.Heading, ins.N/90)
		}
	}
}

func (s *Ship) DistanceFromStart() int {
	return int(math.Abs(float64(s.NS)) + math.Abs(float64(s.EW)))
}

func (s *Ship) Move(dir Direction, num int) {
	switch dir {
	case N:
		s.NS += num
	case S:
		s.NS -= num
	case E:
		s.EW += num
	case W:
		s.EW -= num
	}
}

func turnLeft(dir Direction, times int) Direction {
	for i := 0; i < times; i++ {
		switch dir {
		case N:
			dir = W
		case W:
			dir = S
		case S:
			dir = E
		case E:
			dir = N
		}
	}
	return dir
}

func turnRight(dir Direction, times int) Direction {
	for i := 0; i < times; i++ {
		switch dir {
		case N:
			dir = E
		case E:
			dir = S
		case S:
			dir = W
		case W:
			dir = N
		}
	}
	return dir
}
