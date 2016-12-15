package gobot

import (
	"errors"
	"regexp"
	"strconv"
)

func Parse(input string) (Executable, error) {
	move := regexp.MustCompile(`^\s*MOVE\s*$`)
	place := regexp.MustCompile(`^\s*PLACE\s(\d),(\d),(SOUTH)$`)
	switch {
	case move.MatchString(input):
		return *new(Move), nil
	case place.MatchString(input):
		var groups []string = place.FindStringSubmatch(input)
		return placeFromString(groups[1:]), nil
	default:
		//TODO put the command into a message
		return *new(Move), errors.New("unknown command")
	}
}

func placeFromString(groups []string) Place {
	x, _ := strconv.Atoi(groups[0])
	y, _ := strconv.Atoi(groups[1])
	var facing Direction
	switch groups[2] {
	case "SOUTH":
		facing = SOUTH
	}
	return Place{x, y, facing}
}
