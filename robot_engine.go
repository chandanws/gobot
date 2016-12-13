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
	Table *Table
	X, Y  int
}

func (err outOfBoundsError) Error() string {
	return fmt.Sprintf("Placing robot out of table (%v) x: %d y: %d", err.Table, err.X, err.Y)
}

type Table struct {
	X, Y int
}

type Robot struct {
	X, Y   int
	Facing Direction
}

func Run(table *Table, command Command, x int, y int, direction Direction) (Robot, error) {
	if x >= table.X || y >= table.Y {
		return *new(Robot), outOfBoundsError{table, x, y}
	}
	return Robot{x, y, direction}, nil
}
