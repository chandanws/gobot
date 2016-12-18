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
		command, _ := Parse(scanner.Text())
		fmt.Println(command)
		var err error
		table, output, err = command.Execute(table)
		fmt.Printf("executed table %+v", table)
		if output != nil {
			fmt.Printf("executed output %s\n", *output)
			writeOutput(bufWriter, output)
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	bufWriter.Flush()
	return nil
}

func writeOutput(out *bufio.Writer, output *string) {
	_, err := out.WriteString(*output)
	if err != nil {
		panic("writing of output failed")
	}
	output = nil
}
