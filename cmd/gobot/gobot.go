package main

import (
	"github.com/viktomas/gobot"
	"os"
)

func main() {
	gobot.RunEngine(os.Stdin, os.Stdout)
}
