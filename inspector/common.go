package inspector

import (
	"fmt"
	"go/ast"
)

// getAssignExpr scans right-hands of assignments and returns the expression that is assigned to the given object.
// Note that multiple vars might be assigned in a single assignment.
func (c *WithStackChecker) getAssignExpr(obj *ast.Object) ast.Expr {
	switch decl := obj.Decl.(type) {
	case *ast.AssignStmt:
		switch len(decl.Rhs) {
		case 0:
			return nil
		case 1:
			return decl.Rhs[0]
		}

		if len(decl.Lhs) != len(decl.Rhs) {
			panic("Unmatched length of lhs and rhs")
		}

		for i, expr := range decl.Lhs {
			if ident, ok := expr.(*ast.Ident); ok && ident.Obj == obj {
				return decl.Rhs[i]
			}
		}
	default:
		panic(fmt.Sprintf("Unimplemented type: %T", decl))
	}

	return nil
}
