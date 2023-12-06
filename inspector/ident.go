package inspector

import (
	"go/ast"
)

// checkIdent checks the identifier and report.
func (c *WithStackChecker) checkIdent(ident *ast.Ident) {
	expr := c.getAssignExprInObject(ident.Obj)
	c.checkExpr(expr)
}
