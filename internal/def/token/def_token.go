package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	TYPE      = "TYPE"
	ASSIGN    = "="
	LBRACE    = "{"
	RBRACE    = "{"
)

var Types string = `u8,
	u16,
	u32,
	[]u8,
	i32,`


func LookupIdent(ident string) TokenType {
	if strings.Contains(Types, ident) {
		return TYPE
	} else {
		return IDENT
	}
}

