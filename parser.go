package gobot

import (
	"fmt"
	"regexp"
	"strconv"
)

func Parse(input string) (Executable, error) {
	move := regexp.MustCompile(`^\s*MOVE\s*$`)
	place := regexp.MustCompile(`^\s*PLACE\s(\d+),(\d+),(NORTH|EAST|SOUTH|WEST)$`)
	switch {
	case move.MatchString(input):
		return *new(Move), nil
	case place.MatchString(input):
		var groups []string = place.FindStringSubmatch(input)
		return placeFromString(groups[1:]), nil
	default:
		return *new(Move), fmt.Errorf("Unknown command %s", input)
	}
}

func placeFromString(groups []string) Place {
	x, _ := strconv.Atoi(groups[0])
	y, _ := strconv.Atoi(groups[1])
	var facing Direction
	switch groups[2] {
	case "NORTH":
		facing = NORTH
	case "EAST":
		facing = EAST
	case "SOUTH":
		facing = SOUTH
	case "WEST":
		facing = WEST
	}
	return Place{x, y, facing}
}
