package gobot

import "testing"

func TestPlace(t *testing.T) {
	table := Table{5, 5, *new(Robot), false}
	robot, _ := Run(table, PLACE, 1, 2, SOUTH)
	expectedRobot := Robot{1, 2, SOUTH}
	if robot != expectedRobot {
		t.Errorf("robot %+v is not equal %+v", robot, expectedRobot)
	}
}

func TestPlaceOutOfBounds(t *testing.T) {
	table := Table{5, 5, *new(Robot), false}
	_, err := Run(table, PLACE, 5, 1, EAST)
	if err == nil {
		t.Errorf("putting robot out of bounds didn't cause error")
	}
}

/*
func TestMove(t *testing.T) {
	table := Table{5, 5}
	robot := Rorbot{1, 2, SOUTH}
	newRobot := Run(&table, MOVE, robot)
	expectedRobot := Robot{1, 1, SOUTH}
	if newRobot != expectedRobot {
		t.Errorf("robot %+v\n is not Robot{1,1,SOUTH}", robot)
	}
}
*/
