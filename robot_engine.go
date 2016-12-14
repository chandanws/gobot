package gobot

import "fmt"
import "errors"

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

func Run(table Table, command Command, x int, y int, direction Direction) (Table, error) {
	switch command {
	case PLACE:
		if x >= table.width || y >= table.height {
			return *new(Table), outOfBoundsError{table, x, y}
		}
		return Table{table.height, table.width, Robot{x, y, direction}, true}, nil
	case MOVE:
		movedX, movedY := move(table.robot.x, table.robot.y, table.robot.facing)
		return Table{table.height, table.width, Robot{movedX, movedY, table.robot.facing}, true}, nil

	default:
		return *new(Table), errors.New("unrecognized cmmand")
	}

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
