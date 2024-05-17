package printer

import (
	"fmt"

	"github.com/tredstart/laml/internal/def/ast"
)

func WalkObjects(obj ast.Object) string {
    result := ""
	for _, s := range obj.Statements {
		switch s.Type() {
		case "Var":
			// skip ok check cuz we know the type
			stmt, _ := s.(*ast.VarStatement)
			if stmt.Name == nil {
				result += fmt.Sprintf("pub const %s = struct {\n", stmt.TokenLiteral())
				for _, f := range stmt.Value.Statements {
					if field, ok := f.(*ast.DefStatement); ok {
                        result += fmt.Sprintf("    %s: %s,\n", field.Value.Value, field.TokenLiteral())
					} else {
						fmt.Println("something really really wrong went here")
						return ""
					}
				}
				result += fmt.Sprintln("};")
			} else {
				result += fmt.Sprintf("pub var %s = %s{\n", stmt.Name, stmt.TokenLiteral())
				for _, f := range stmt.Value.Statements {
					if field, ok := f.(*ast.FieldStatement); ok {
						result += fmt.Sprintf("    .%s = %s,\n", field.TokenLiteral(), field.Value.Value)
					} else {
						fmt.Println("something really really wrong went here")
						return ""
					}
				}
				result += fmt.Sprintln("};")
			}
		default:
			fmt.Println("something really really wrong went here")
			return ""
		}
	}
    return result
}
