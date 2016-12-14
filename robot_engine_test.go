package gobot

import "testing"

func TestTableContains(t *testing.T) {
	table := Table{5, 5, *new(Robot), false}
	if !table.contains(3, 3) {
		t.Error("Table should contain 3,3")
	}
	if table.contains(-1, 0) {
		t.Error("Table shouldn't contain -1,0")
	}
	if table.contains(5, 0) {
		t.Error("Table shouldn't contain 5,0")
	}
	if table.contains(0, -1) {
		t.Error("Table shouldn't contain 0,-1")
	}
	if table.contains(0, 5) {
		t.Error("Table shouldn't contain 0,5")
	}
}
