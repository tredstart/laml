package parser

import (
	"fmt"
	"littlelang/internal/def/ast"
	"littlelang/internal/def/lexer"
	"littlelang/internal/def/token"
	"reflect"
)

type Parser struct {
	l *lexer.Lexer

	errors    []string
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.NextToken()
	p.NextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) PeekErrors(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseObjects() *ast.Object {
	program := &ast.Object{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		s := p.parseStatement()
		if s != nil && !reflect.ValueOf(s).IsNil() {
			program.Statements = append(program.Statements, s)
		}
		p.NextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.IDENT:
		return p.parseVarStatement()
	default:
		return nil
	}
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	s := &ast.VarStatement{Token: p.curToken}
	if !p.peekTokenIs(token.COLON) {
		if !p.expectPeek(token.IDENT) {
			p.PeekErrors(token.IDENT)
			return nil
		}
		s.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		if !p.expectPeek(token.COLON) {
			p.PeekErrors(token.COLON)
			return nil
		}
		for !p.peekTokenIs(token.SEMICOLON){

			if p.curTokenIs(token.EOF) {
				p.PeekErrors(token.SEMICOLON)
				return nil
			}
			obj := p.parseFieldStatement()

			if obj == nil {
				return nil
			}
			s.Value.Statements = append(s.Value.Statements, obj)
		}
	} else {
		// TODO make this thing def statement
	}
	return s
}

func (p *Parser) parseFieldStatement() *ast.FieldStatement {

	if !p.expectPeek(token.IDENT) {
		p.PeekErrors(token.IDENT)
		return nil
	}

	s := &ast.FieldStatement{
		Token: token.Token{Type: token.IDENT, Literal: p.curToken.Literal},
	}

	if !p.expectPeek(token.ASSIGN) {
		p.PeekErrors(token.ASSIGN)
		return nil
	}

	if !p.expectPeek(token.IDENT) {
		p.PeekErrors(token.IDENT)
		return nil
	}

	s.Value = &ast.Identifier{
		Token: token.Token{Type: token.IDENT, Literal: p.curToken.Literal},
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.COMMA) {
		p.PeekErrors(token.COMMA)
		return nil
	}

	return s

}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	} else {
		p.PeekErrors(t)
		return false
	}
}
