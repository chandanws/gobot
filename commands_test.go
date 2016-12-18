package gobot

import "testing"

func TestPlaceCommand(t *testing.T) {
	table := Table{5, 5, *new(Robot), false}
	place := Place{1, 2, SOUTH}
	newTable, _, _ := place.Execute(table)
	expectedTable := Table{5, 5, Robot{1, 2, SOUTH}, true}
	if newTable != expectedTable {
		t.Errorf("table %+v is not equal %+v", newTable, expectedTable)
	}
}

func TestPlaceOutOfBounds(t *testing.T) {
	table := Table{5, 5, *new(Robot), false}
	place := Place{5, 1, EAST}
	_, _, err := place.Execute(table)
	if err == nil {
		t.Errorf("putting robot out of bounds didn't cause error")
	}
}

func TestPlaceOutOfBounds2(t *testing.T) {
	table := Table{5, 5, Robot{1, 1, SOUTH}, false}
	place := Place{4, -1, EAST}
	newTable, _, err := place.Execute(table)
	if err == nil {
		t.Error("putting robot out of bounds didn't cause error")
	}
	if newTable != table {
		t.Error("unsucessful place command changed table")
	}
}

func TestReport(t *testing.T) {
	table := Table{5, 5, Robot{4, 1, NORTH}, true}
	report := *new(Report)
	newTable, stdio, _ := report.Execute(table)
	if *stdio != "4,1,NORTH" {
		t.Errorf("Reported %s when expecting 4,1,NORTH", *stdio)
	}
	if newTable != table {
		t.Error("unsucessful report command changed table")
	}
}

func TestReportOnUninitializedTable(t *testing.T) {
	report := *new(Report)
	testUninitialized(report, "report", t)
}

func TestMoveCommand(t *testing.T) {
	table := Table{5, 5, Robot{1, 2, SOUTH}, true}
	move := *new(Move)
	newTable, _, _ := move.Execute(table)
	expectedTable := Table{5, 5, Robot{1, 1, SOUTH}, true}
	if newTable != expectedTable {
		t.Errorf("table %+v is not equal %+v", newTable, expectedTable)
	}
}

func TestMoveCommandOutOfBounds(t *testing.T) {
	table := Table{5, 5, Robot{1, 0, SOUTH}, true}
	move := *new(Move)
	newTable, _, err := move.Execute(table)
	if err == nil {
		t.Error("moving robot out of bounds didn't cause error")
	}
	if newTable != table {
		t.Error("unsucessful report command changed table")
	}
}

func TestMoveOnUninitializedTable(t *testing.T) {
	move := *new(Move)
	testUninitialized(move, "move", t)
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

func TestLeft(t *testing.T) {
	left := Left{}
	table := Table{5, 5, Robot{1, 2, SOUTH}, true}
	expected := Table{5, 5, Robot{1, 2, EAST}, true}
	testCommand(left, "left", table, expected, t)
}

func TestLeftUninitialized(t *testing.T) {
	left := Left{}
	testUninitialized(left, "left", t)
}

func testCommand(e Executable, name string, table Table, expected Table, t *testing.T) {
	newTable, _, _ := e.Execute(table)
	if newTable != expected {
		t.Errorf(
			"Command %s%v failed, it transformed table %v to %v instead %v",
			name,
			e,
			table,
			newTable,
			expected,
		)
	}
}

func testUninitialized(e Executable, name string, t *testing.T) {
	table := Table{5, 5, Robot{1, 2, SOUTH}, false}
	_, _, err := e.Execute(table)
	if err == nil {
		t.Errorf(
			"Command %s%v failed, it was sucessfuly run on uninitialized table",
			name,
			e,
		)
	}
}
