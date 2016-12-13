package gobot

type Command int
type Direction int

const (
	PLACE Command = iota
)

const (
	NORTH Direction = iota
	EAST  Direction = iota
	SOUTH Direction = iota
	WEST  Direction = iota
)

type Table struct {
	X, Y int
}

type Robot struct {
	X, Y   int
	Facing Direction
}

func Run(table *Table, command Command, x int, y int, direction Direction) Robot {
	return Robot{x, y, direction}
}
