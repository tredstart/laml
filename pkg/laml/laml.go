package laml

import (
	"fmt"

	"github.com/tredstart/laml/internal/def/ast"
	"github.com/tredstart/laml/internal/def/lexer"
	"github.com/tredstart/laml/internal/def/parser"
	"github.com/tredstart/laml/internal/printer"
)

func LamlParse(input string) *ast.Object {
	l := lexer.New(string(input))
	p := parser.New(l)

	result := p.ParseObjects()
	for _, e := range p.Errors() {
		fmt.Printf("e: %v\n", e)
	}
	return result
}

func LamlWalk(obj ast.Object) string {
	return printer.WalkObjects(obj)
}
