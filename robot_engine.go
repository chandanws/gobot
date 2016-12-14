package gobot

import "fmt"

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

type outOfBoundsError struct {
	Table Table
	x, y  int
}

func (err outOfBoundsError) Error() string {
	return fmt.Sprintf("Placing robot out of table (%v) x: %d y: %d", err.Table, err.x, err.y)
}

type Table struct {
	width, height int
	robot         Robot
	initialized   bool
}

type Robot struct {
	x, y   int
	facing Direction
}

func Run(table Table, command Command, x int, y int, direction Direction) (Robot, error) {
	if x >= table.width || y >= table.height {
		return *new(Robot), outOfBoundsError{table, x, y}
	}
	return Robot{x, y, direction}, nil
}
