package lexer

import (
	"laml/internal/def/token"
	"testing"
)

func TestNextToken(t *testing.T) {
    input := `
velocity:
    i32 x,
    u8 y,
;

position: u32 x, i32 y;
box:
    u32 width,
    u32 height, 
    position pos,
;

cmp.player_def player:
    position = default_pos,
    velocity = {x=0, y=70},
;
    `
    tests := []struct{
        expectedType token.TokenType
        expectedLiteral string
    }{
        {token.IDENT, "velocity"},
        {token.COLON, ":"},
        {token.IDENT, "i32"},
        {token.IDENT, "x"},
        {token.COMMA, ","},
        {token.IDENT, "u8"},
        {token.IDENT, "y"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},

        {token.IDENT, "position"},
        {token.COLON, ":"},
        {token.IDENT, "u32"},
        {token.IDENT, "x"},
        {token.COMMA, ","},
        {token.IDENT, "i32"},
        {token.IDENT, "y"},
        {token.SEMICOLON, ";"},

        {token.IDENT, "box"},
        {token.COLON, ":"},
        {token.IDENT, "u32"},
        {token.IDENT, "width"},
        {token.COMMA, ","},
        {token.IDENT, "u32"},
        {token.IDENT, "height"},
        {token.COMMA, ","},
        {token.IDENT, "position"},
        {token.IDENT, "pos"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},

        {token.IDENT, "cmp.player_def"},
        {token.IDENT, "player"},
        {token.COLON, ":"},
        {token.IDENT, "position"},
        {token.ASSIGN, "="},
        {token.IDENT, "default_pos"},
        {token.COMMA, ","},
        {token.IDENT, "velocity"},
        {token.ASSIGN, "="},
        {token.LBRACE, "{"},
        {token.IDENT, "x"},
        {token.ASSIGN, "="},
        {token.IDENT, "0"},
        {token.COMMA, ","},
        {token.IDENT, "y"},
        {token.ASSIGN, "="},
        {token.IDENT, "70"},
        {token.RBRACE, "}"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},

    }
    l := New(input)
    for i, tt := range tests {
        tok := l.NextToken()
        if tok.Type != tt.expectedType {
            t.Log(l.readPosition)
            t.Log(l.input[l.readPosition:])
            t.Fatalf("tests[%d] - def tokentype wrong expected %q/%q got %q",
                i, tt.expectedType, tt.expectedLiteral, tok.Type)
        }
        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - def tokentype wrong expected %q got %q",
                i, tt.expectedLiteral, tok.Literal)
        }
    }
}
