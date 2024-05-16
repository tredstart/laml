package token


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
	ASSIGN    = "="
	LBRACE    = "{"
	RBRACE    = "}"
)

