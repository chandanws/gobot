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

type Robot struct {
	x, y   int
	facing Direction
}
