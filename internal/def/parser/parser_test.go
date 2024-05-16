package parser

import (
	"littlelang/internal/def/ast"
	"littlelang/internal/def/lexer"
	"littlelang/internal/def/token"
	"testing"
)

func TestVarStatements(t *testing.T) {
	input := `
    cmp.position default_pos:
    x = 50,
    y = 0,
;
    `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseObjects()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("ParseObjects failed to init")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements doesn't have 1 statements. got = %d", len(program.Statements))
	}

	tests := []struct {
		expectedType       string
		expectedIdentifier string
	}{
		{"cmp.position", "default_pos"},
	}

	for i, tt := range tests {
		stms := program.Statements[i]
        t.Log(program.Statements)
		if !testVarStatement(t, stms, tt.expectedType, tt.expectedIdentifier) {
			t.Fatal("Failed to get correct statemetn")
		}
	}
}

func testVarStatement(t *testing.T, s ast.Statement, tt, name string) bool {
	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement. got = %T", s)
		return false
	}

	if varStmt.TokenLiteral() != tt {
		t.Errorf("letStmt.Type.Value not %s. got = %s ", tt, varStmt.TokenLiteral())
		return false
	}

	if varStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. got = %s ", name, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral not %s. got = %s ", name, varStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser returned with %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parse error: %q", msg)
	}
	t.FailNow()
}

func TestString(t *testing.T) {
	prog := ast.Object{
		Statements: []ast.Statement{
			&ast.VarStatement{
				Token: token.Token{Type: token.IDENT, Literal: "position"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "default_pos"},
					Value: "default_pos",
				},
				Value: ast.Object{
					Statements: []ast.Statement{
						&ast.FieldStatement{
                            Token: token.Token{Type: token.IDENT, Literal: "x"},
                            Value: &ast.Identifier{
                                Token: token.Token{Type: token.IDENT, Literal: "0"},
                                Value: "0",
                            },
                        },
						&ast.FieldStatement{
                            Token: token.Token{Type: token.IDENT, Literal: "y"},
                            Value: &ast.Identifier{
                                Token: token.Token{Type: token.IDENT, Literal: "70"},
                                Value: "70",
                            },
                        },
					},
				},
			},
		},
	}

	if prog.String() != "position default_pos: x = 0,\ny = 70,\n;" {
		t.Errorf("somethign worng : %q", prog.String())
	}
}
