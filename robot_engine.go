package gobot

import (
	"bufio"
	"fmt"
	"io"
)

func RunEngine(in io.Reader, out io.Writer) error {
	bufWriter := bufio.NewWriter(out)
	scanner := bufio.NewScanner(in)
	table := Table{5, 5, *new(Robot), false}
	var output *string
	for scanner.Scan() {
		var err error
		command, err := Parse(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		table, output, err = command.Execute(table)
		if output != nil {
			writeOutput(bufWriter, output)
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func writeOutput(out *bufio.Writer, output *string) {
	_, err := out.WriteString(*output)
	if err != nil {
		panic("writing of output failed")
	}
	output = nil
	out.Flush()
}
