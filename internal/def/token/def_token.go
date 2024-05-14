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
    TYPE = "TYPE"
)

var types = []string{
	"u8",
	"u16",
	"u32",
	"[]u8",
    "i32",
}

func contains(str string) bool {
    for _, v := range types {
        if v == str {
            return true
        }
    }
    return false
}

func LookupIdent(ident string) TokenType {
    if contains(ident) {
        return TYPE
    } else {
        return IDENT
    }
}
