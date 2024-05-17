package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tredstart/laml/def/lexer"
	"github.com/tredstart/laml/def/parser"
	"github.com/tredstart/laml/printer"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("expected a filename as an argument")
		return
	}
	filename := args[1]

	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("cannot read file: " + err.Error())
		return
	}

    l := lexer.New(string(input))
    p := parser.New(l)

	program := p.ParseObjects()
    for _, e := range p.Errors() {
        fmt.Println(e)
    }
	if program == nil {
		fmt.Println("can't parse")
		return
	}
	result := printer.WalkObjects(*program)
	filename = strings.TrimSuffix(filename, ".laml")
	filename += ".zig"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("cannot create file")
		return
	}
	_, err = file.WriteString(result)
	if err != nil {
		fmt.Println("cannot write file")
		return
	}
}
