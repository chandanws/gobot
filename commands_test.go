package gobot

import "testing"

func TestPlaceCommand(t *testing.T) {
	table := Table{5, 5, *new(Robot), false}
	place := Place{1, 2, SOUTH}
	newTable, _ := place.Execute(table)
	expectedTable := Table{5, 5, Robot{1, 2, SOUTH}, true}
	if newTable != expectedTable {
		t.Errorf("table %+v is not equal %+v", newTable, expectedTable)
	}
}

func TestPlaceOutOfBounds(t *testing.T) {
	table := Table{5, 5, *new(Robot), false}
	place := Place{5, 1, EAST}
	_, err := place.Execute(table)
	if err == nil {
		t.Errorf("putting robot out of bounds didn't cause error")
	}
}

func TestMoveCommand(t *testing.T) {
	table := Table{5, 5, Robot{1, 2, SOUTH}, true}
	move := *new(Move)
	newTable, _ := move.Execute(table)
	expectedTable := Table{5, 5, Robot{1, 1, SOUTH}, true}
	if newTable != expectedTable {
		t.Errorf("table %+v is not equal %+v", newTable, expectedTable)
	}
}

func TestMoveNorth(t *testing.T) {
	x, y := move(0, 0, NORTH)
	if x != 0 || y != 1 {
		t.Errorf("moved to %d,%d instad of 0,1", x, y)
	}
}

func TestMoveEast(t *testing.T) {
	x, y := move(1, 1, EAST)
	if x != 2 || y != 1 {
		t.Errorf("moved to %d,%d instad of 0,1", x, y)
	}
}

func TestMoveSouth(t *testing.T) {
	x, y := move(2, 2, SOUTH)
	if x != 2 || y != 1 {
		t.Errorf("moved to %d,%d instad of 0,1", x, y)
	}
}

func TestMoveWest(t *testing.T) {
	x, y := move(1, 1, WEST)
	if x != 0 || y != 1 {
		t.Errorf("moved to %d,%d instad of 0,1", x, y)
	}
}
