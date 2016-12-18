package gobot

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestE2E(t *testing.T) {
	reader := strings.NewReader("PLACE 0,0,NORTH\nREPORT")
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	RunEngine(reader, writer)
	result := buffer.String()
	if result != "0,0,NORTH" {
		t.Errorf("Engine returned invalid result (%s) instead of 0,0,NORTH", result)
	}
}
