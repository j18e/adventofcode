package main

type Grid struct {
	Rows, Seats int
	Spots       map[int]Spot
}

func (g *Grid) FillSeats() {
	var toFill, toEmpty []int
	for id, spot := range g.Spots {
		if spot == SpotFloor {
			continue
		}
		adjOcc := g.CountAdjacentOccupied(id)
		if spot == SpotEmpty && adjOcc == 0 {
			toFill = append(toFill, id)
		}
		if spot == SpotOccupied && adjOcc >= 4 {
			toEmpty = append(toEmpty, id)
		}
	}
	for _, id := range toFill {
		g.Spots[id] = SpotOccupied
	}
	for _, id := range toEmpty {
		g.Spots[id] = SpotEmpty
	}
}

func (g *Grid) FillSeats2() {
	var toFill, toEmpty []int
	for id, spot := range g.Spots {
		if spot == SpotFloor {
			continue
		}
		occ := g.CountOccupiedAllDirections(id)
		if spot == SpotEmpty && occ == 0 {
			toFill = append(toFill, id)
		}
		if spot == SpotOccupied && occ >= 5 {
			toEmpty = append(toEmpty, id)
		}
	}
	for _, id := range toFill {
		g.Spots[id] = SpotOccupied
	}
	for _, id := range toEmpty {
		g.Spots[id] = SpotEmpty
	}
}

func (g *Grid) TotalOccupied() int {
	var cnt int
	for _, s := range g.Spots {
		if s == SpotOccupied {
			cnt++
		}
	}
	return cnt
}

func (g *Grid) CountOccupiedAllDirections(id int) int {
	return g.occupiedInDirections(g.coordFromID(id), N, NE, E, SE, S, SW, W, NW)
}

func (g *Grid) occupiedInDirections(c coordinate, dx ...Direction) int {
	var cnt int
	if g.FirstSeatInDirection(c, dx[0]) == SpotOccupied {
		cnt++
	}
	if len(dx) == 1 {
		return cnt
	}
	return cnt + g.occupiedInDirections(c, dx[1:]...)
}

func (g *Grid) FirstSeatInDirection(c coordinate, dir Direction) Spot {
	if dir == N || dir == NE || dir == NW {
		c.row--
	}
	if dir == S || dir == SE || dir == SW {
		c.row++
	}
	if dir == E || dir == NE || dir == SE {
		c.seat++
	}
	if dir == W || dir == NW || dir == SW {
		c.seat--
	}
	spot := g.SpotFromCoord(c)
	if spot == SpotFloor {
		return g.FirstSeatInDirection(c, dir)
	}
	return spot
}

func (g *Grid) CountAdjacentOccupied(id int) int {
	row, seat := g.SeatFromID(id)
	return g.SumOccupied([]coordinate{
		{row - 1, seat - 1}, {row - 1, seat}, {row - 1, seat + 1},
		{row, seat - 1}, {row, seat + 1},
		{row + 1, seat - 1}, {row + 1, seat}, {row + 1, seat + 1},
	})
}

type coordinate struct {
	row, seat int
}

func (g *Grid) SumOccupied(cx []coordinate) int {
	var cnt int
	if g.seatOccupied(cx[0]) {
		cnt++
	}
	if len(cx) == 1 {
		return cnt
	}
	return cnt + g.SumOccupied(cx[1:])
}

func (g *Grid) seatOccupied(c coordinate) bool {
	return g.Spots[g.SeatID(c)] == SpotOccupied
}

func (g *Grid) SeatID(c coordinate) int {
	if c.row < 0 || c.row >= g.Rows || c.seat < 0 || c.seat >= g.Seats {
		return -1
	}
	return c.row*g.Seats + c.seat
}

func (g *Grid) SpotFromCoord(c coordinate) Spot {
	id := g.SeatID(c)
	if id == -1 {
		return SpotOffGrid
	}
	return g.Spots[id]
}

func (g *Grid) coordFromID(id int) coordinate {
	row, seat := g.SeatFromID(id)
	return coordinate{row, seat}
}

func (g *Grid) SeatFromID(id int) (row, seat int) {
	if id < 0 || id > g.Rows*g.Seats {
		return -1, -1
	}
	seat = id % g.Seats
	row = (id - seat) / g.Seats
	return row, seat
}
