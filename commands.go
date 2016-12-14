package gobot

import (
	"errors"
	"fmt"
)

type Executable interface {
	Execute(table Table) (Table, *string, error)
}

type Place struct {
	x, y   int
	facing Direction
}

type Move struct {
}

type outOfBoundsError struct {
	Table Table
	x, y  int
}

func (err outOfBoundsError) Error() string {
	return fmt.Sprintf("Placing robot out of table (%v) x: %d y: %d", err.Table, err.x, err.y)
}

func (place Place) Execute(table Table) (Table, *string, error) {
	if !table.contains(place.x, place.y) {
		return *new(Table), nil, outOfBoundsError{table, place.x, place.y}
	}
	return Table{table.height, table.width, Robot{place.x, place.y, place.facing}, true}, nil, nil
}

func (moveCmd Move) Execute(table Table) (Table, *string, error) {
	if !table.initialized {
		return *new(Table), nil, errors.New("Executing move on uninitialized table")
	}
	x, y := move(table.robot.x, table.robot.y, table.robot.facing)
	if !table.contains(x, y) {
		return *new(Table), nil, outOfBoundsError{table, x, y}
	}
	return Table{table.height, table.width, Robot{x, y, table.robot.facing}, true}, nil, nil
}

func move(x int, y int, facing Direction) (int, int) {
	switch facing {
	case NORTH:
		return x, y + 1
	case EAST:
		return x + 1, y
	case SOUTH:
		return x, y - 1
	case WEST:
		return x - 1, y
	default:
		panic("there is only 4 world directions")
	}
}
