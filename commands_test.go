package gobot

import "testing"

type testCase struct {
	e               Executable
	name            string
	table, expected Table
}

func TestPlaceCommand(t *testing.T) {
	tc := testCase{
		Place{1, 2, SOUTH},
		"place",
		Table{5, 5, *new(Robot), false},
		Table{5, 5, Robot{1, 2, SOUTH}, true},
	}
	testCommands(t, tc)
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
	move := *new(Move)
	testCommands(t,
		testCase{move, "move",
			Table{5, 5, Robot{1, 2, SOUTH}, true},
			Table{5, 5, Robot{1, 1, SOUTH}, true},
		},
		testCase{move, "move",
			Table{5, 5, Robot{0, 0, NORTH}, true},
			Table{5, 5, Robot{0, 1, NORTH}, true},
		},
		testCase{move, "move",
			Table{5, 5, Robot{1, 1, EAST}, true},
			Table{5, 5, Robot{2, 1, EAST}, true},
		},
		testCase{move, "move",
			Table{5, 5, Robot{1, 1, WEST}, true},
			Table{5, 5, Robot{0, 1, WEST}, true},
		},
	)
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

func TestLeft(t *testing.T) {
	testCommands(t,
		testCase{
			Left{}, "left",
			Table{5, 5, Robot{1, 2, SOUTH}, true},
			Table{5, 5, Robot{1, 2, EAST}, true},
		},
		testCase{
			Left{}, "left",
			Table{5, 5, Robot{1, 2, EAST}, true},
			Table{5, 5, Robot{1, 2, NORTH}, true},
		},
		testCase{
			Left{}, "left",
			Table{5, 5, Robot{1, 2, NORTH}, true},
			Table{5, 5, Robot{1, 2, WEST}, true},
		},
		testCase{
			Left{}, "left",
			Table{5, 5, Robot{1, 2, WEST}, true},
			Table{5, 5, Robot{1, 2, SOUTH}, true},
		},
	)
}

func TestLeftUninitialized(t *testing.T) {
	left := Left{}
	testUninitialized(left, "left", t)
}

func testCommands(t *testing.T, testCases ...testCase) {
	for _, tc := range testCases {
		newTable, _, _ := tc.e.Execute(tc.table)
		if newTable != tc.expected {
			t.Errorf(
				"Command %s%v failed, it transformed table %v to %v instead %v",
				tc.name,
				tc.e,
				tc.table,
				newTable,
				tc.expected,
			)
		}
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
