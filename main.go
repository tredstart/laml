package main

import (
	"fmt"
	"littlelang/internal/def/lexer"
	"littlelang/internal/def/parser"
	"littlelang/internal/printer"
)

func main() {
    to_translate := `
velocity:
    i32 x,
    u8 y,
;
cmp.position default_pos:
    x = 50,
    y = 50,
;
    `
	l := lexer.New(to_translate)
	p := parser.New(l)

	program := p.ParseObjects()
	if program == nil {
		fmt.Println("can't parse")
        return
	}
    printer.WalkObjects(*program)
}
