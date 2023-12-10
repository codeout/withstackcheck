package inspector

import (
	"go/ast"
)

// checkIdent checks the identifier and report.
func (c *WithStackChecker) checkIdent(ident *ast.Ident) {
	switch decl := ident.Obj.Decl.(type) {
	case *ast.ValueSpec:
		c.checkValueSpec(decl)
	case *ast.AssignStmt:
		expr := c.findAssignExprInFunction(ident.Obj.Decl)
		c.checkExpr(expr)
	case *ast.Field: // named return
		expr := c.findAssignExprToNamedReturnInFunction(ident)
		c.checkExpr(expr)
	default:
		c.panicf("Unimplemented type: %T at %s", ident.Obj.Decl, c.pass.Fset.Position(c.pos))
	}
}
