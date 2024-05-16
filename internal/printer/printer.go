package printer

import (
	"fmt"
	"littlelang/internal/def/ast"
)

func WalkObjects(obj ast.Object) {
	for _, s := range obj.Statements {
		switch s.Type() {
		case "Var":
			// skip ok check cuz we know the type
			stmt, _ := s.(*ast.VarStatement)
			if stmt.Name == nil {
				fmt.Printf("pub const %s = struct {\n", stmt.TokenLiteral())
				for _, f := range stmt.Value.Statements {
					if field, ok := f.(*ast.DefStatement); ok {
						fmt.Printf("    %s %s,\n", field.Value.Value, field.TokenLiteral())
					} else {
						fmt.Println("something really really wrong went here")
						return
					}
				}
				fmt.Println("};")
			} else {
				fmt.Printf("pub var %s = %s{\n", stmt.Name, stmt.TokenLiteral())
				for _, f := range stmt.Value.Statements {
					if field, ok := f.(*ast.FieldStatement); ok {
						fmt.Printf("    .%s = %s,\n", field.TokenLiteral(), field.Value.Value)
					} else {
						fmt.Println("something really really wrong went here")
						return
					}
				}
				fmt.Println("};")
			}
		default:
			fmt.Println("something really really wrong went here")
			return
		}
	}
}
