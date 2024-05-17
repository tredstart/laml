package printer

import (
	"laml/internal/def/lexer"
	"laml/internal/def/parser"
	"testing"
)

func TestWrite(t *testing.T) {
	input := `
velocity:
    i32 x,
    u8 y,
;
position default_pos:
    x = 50,
    y = 70,
;
    `
	expected := `pub const velocity = struct {
    x i32,
    y u8,
};
pub var default_pos = position{
    .x = 50,
    .y = 70,
};
`
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseObjects()
	if program == nil {
		t.Errorf("can't parse %q", program)
	}
	result := WalkObjects(*program)
	if result != expected {
		t.Fatalf("expected \n%s got:\n%s", expected, result)
	}
}
