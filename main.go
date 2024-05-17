package main

import (
	"fmt"
	"laml/internal/def/lexer"
	"laml/internal/def/parser"
	"laml/internal/printer"
	"os"
	"strings"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("expected a filename as an argument")
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
	if program == nil {
		fmt.Println("can't parse")
		return
	}
    result := printer.WalkObjects(*program)
    for _, e := range p.Errors(){
        fmt.Printf("e: %v\n", e)
    }
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
