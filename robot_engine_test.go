package gobot

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestE2E(t *testing.T) {
	testE2E("PLACE 0,0,NORTH\nREPORT", "0,0,NORTH", t)
	testE2E("PLACE 0,0,NORTH\nMOVE\nREPORT", "0,1,NORTH", t)
	testE2E("PLACE 0,0,NORTH\nRUBBISH\nREPORT", "0,0,NORTH", t)
}

func testE2E(input string, output string, t *testing.T) {
	reader := strings.NewReader(input)
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	RunEngine(reader, writer)
	result := buffer.String()
	if result != output {
		t.Errorf("Engine returned invalid result %s instead of %s", result, output)
	}
}
