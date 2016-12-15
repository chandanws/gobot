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

func TestParsingRubbish(t *testing.T) {
	_, err := Parse("Rubbish")
	if err == nil {
		t.Error("Rubbish command didn't cause and error")
	}
}

func TestParsingPlace(t *testing.T) {
	cmd, _ := Parse("PLACE 2,3,SOUTH")
	switch cmd.(type) {
	case Place:
		place := cmd.(Place)
		expected := Place{2, 3, SOUTH}
		if place != expected {
			t.Errorf("The place command wasn't parsed correctly, %+v", place)
		}
	default:
		t.Error("The command wasn't parsed as place")
	}
}
