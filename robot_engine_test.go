package gobot

import "testing"

func TestPlace(t *testing.T) {
	table := Table{5, 5}
	robot := Run(&table, PLACE, 1, 2, SOUTH)
	if robot.X != 1 {
		t.Errorf("robot.X is %d, want %d", robot.X, 1)
	}
	if robot.Y != 2 {
		t.Errorf("robot.Y is %d, want %d", robot.Y, 2)
	}
	if robot.Facing != SOUTH {
		t.Errorf("robot.Facing is %d, want %d", robot.Facing, SOUTH)
	}

}
