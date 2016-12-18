package gobot

type Command int
type Direction int

const (
	PLACE Command = iota
	MOVE  Command = iota
)

const (
	NORTH Direction = iota
	EAST  Direction = iota
	SOUTH Direction = iota
	WEST  Direction = iota
)

type Table struct {
	width, height int
	robot         Robot
	initialized   bool
}

func (table Table) contains(x, y int) bool {
	return x >= 0 && x < table.width &&
		y >= 0 && y < table.height
}

type Robot struct {
	x, y   int
	facing Direction
}
