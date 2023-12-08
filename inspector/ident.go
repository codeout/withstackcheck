package inspector

import (
	"fmt"
	"go/ast"
)

// checkIdent checks the identifier and report.
func (c *WithStackChecker) checkIdent(ident *ast.Ident) {
	switch decl := ident.Obj.Decl.(type) {
	case *ast.ValueSpec:
		// need to find var assignments when "var err error" is found
		c.checkValueSpec(decl)
		return
	case *ast.AssignStmt:
		expr := c.findAssignExprInFunction(ident.Obj.Decl)
		c.checkExpr(expr)
	case *ast.Field: // named return
		expr := c.findAssignExprToNamedReturnInFunction(ident)
		c.checkExpr(expr)
	default:
		panic(fmt.Sprintf("Unimplemented type: %T", ident.Obj.Decl))
	}
}
