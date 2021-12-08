package main

import (
	"fmt"
)

func main() {
	ix := parse()
	part1(ix)
	part2(ix)
}

func part1(ix []Instruction) {
	ship := &Ship{E, 0, 0}
	ship.FollowCourse(ix)
	fmt.Println(ship.DistanceFromStart())
}

func part2(ix []Instruction) {
	ship := &Ship{}
	wp := &Waypoint{1, 10}
	ship.FollowCourseWithWaypoint(wp, ix)
	fmt.Println(ship.DistanceFromStart())
}
