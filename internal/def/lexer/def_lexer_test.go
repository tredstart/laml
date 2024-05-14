package lexer

import (
	"littlelang/internal/def/token"
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
    `
    tests := []struct{
        expectedType token.TokenType
        expectedLiteral string
    }{
        {token.IDENT, "velocity"},
        {token.COLON, ":"},
        {token.TYPE, "i32"},
        {token.IDENT, "x"},
        {token.COMMA, ","},
        {token.TYPE, "u8"},
        {token.IDENT, "y"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},

        {token.IDENT, "position"},
        {token.COLON, ":"},
        {token.TYPE, "u32"},
        {token.IDENT, "x"},
        {token.COMMA, ","},
        {token.TYPE, "i32"},
        {token.IDENT, "y"},
        {token.SEMICOLON, ";"},

        {token.IDENT, "box"},
        {token.COLON, ":"},
        {token.TYPE, "u32"},
        {token.IDENT, "width"},
        {token.COMMA, ","},
        {token.TYPE, "u32"},
        {token.IDENT, "height"},
        {token.COMMA, ","},
        {token.IDENT, "position"},
        {token.IDENT, "pos"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},
    }
    l := New(input)
    for i, tt := range tests {
        tok := l.NextToken()
        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - def tokentype wrong expected %q got %q",
                i, tt.expectedType, tok.Type)
        }
        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - def tokentype wrong expected %q got %q",
                i, tt.expectedLiteral, tok.Literal)
        }
    }
}
