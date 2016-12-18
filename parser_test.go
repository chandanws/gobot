package gobot

import "testing"

func TestParsingMove(t *testing.T) {
	cmd, _ := Parse("MOVE")
	switch cmd.(type) {
	case Move:
		return
	default:
		t.Error("The command wasn't parsed as move")
	}
}

func TestParsingLeft(t *testing.T) {
	cmd, _ := Parse("LEFT")
	switch cmd.(type) {
	case Left:
		return
	default:
		t.Error("The command wasn't parsed as left")
	}
}

func TestParsingRubbish(t *testing.T) {
	_, err := Parse("Rubbish")
	if err == nil {
		t.Error("Rubbish command didn't cause an error")
	}
}

func TestParsingPlace(t *testing.T) {
	testParsePlace("PLACE 10,20,NORTH", Place{10, 20, NORTH}, t)
	testParsePlace("PLACE 0,0,EAST", Place{0, 0, EAST}, t)
	testParsePlace("PLACE 2,3,SOUTH", Place{2, 3, SOUTH}, t)
	testParsePlace("PLACE 7,2,WEST", Place{7, 2, WEST}, t)
}

func TestParsingReport(t *testing.T) {
	cmd, _ := Parse("REPORT")
	switch cmd.(type) {
	case Report:
		return
	default:
		t.Error("REPORT didn't get parsed as command")
	}
}

func testParsePlace(input string, command Executable, t *testing.T) {
	cmd, _ := Parse(input)
	switch cmd.(type) {
	case Place:
		place := cmd.(Place)
		if place != command {
			t.Errorf("The place command wasn't parsed correctly, %+v", place)
		}
	default:
		t.Error("The command wasn't parsed as place")
	}
}
