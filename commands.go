package gobot

import "fmt"

type Executable interface {
	Execute(table Table) (Table, error)
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

func (place Place) Execute(table Table) (Table, error) {
	if place.x >= table.width || place.y >= table.height {
		return *new(Table), outOfBoundsError{table, place.x, place.y}
	}
	return Table{table.height, table.width, Robot{place.x, place.y, place.facing}, true}, nil
}

func (moveCmd Move) Execute(table Table) (Table, error) {
	movedX, movedY := move(table.robot.x, table.robot.y, table.robot.facing)
	return Table{table.height, table.width, Robot{movedX, movedY, table.robot.facing}, true}, nil
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
