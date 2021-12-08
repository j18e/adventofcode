package main

type Waypoint struct {
	NS, EW int
}

func (w *Waypoint) Move(dir Direction, dst int) {
	switch dir {
	case N:
		w.NS += dst
	case S:
		w.NS -= dst
	case E:
		w.EW += dst
	case W:
		w.EW -= dst
	default:
		panic("direction not N S E W")
	}
}

func (w *Waypoint) Turn(dir Direction, times int) {
	var fn func()
	switch dir {
	case L:
		fn = w.turnLeft
	case R:
		fn = w.turnRight
	default:
		panic("direction not L or R")
	}
	for i := 0; i < times; i++ {
		fn()
	}
}

func (w *Waypoint) turnRight() {
	w.NS, w.EW = w.EW*-1, w.NS
}

func (w *Waypoint) turnLeft() {
	w.NS, w.EW = w.EW, w.NS*-1
}
