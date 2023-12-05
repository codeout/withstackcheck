package inspector

import (
	"go/ast"
	"go/token"
)

// checkIdent checks the identifier and report.
func (c *WithStackChecker) checkIdent(ident *ast.Ident, pos token.Pos) {
	expr := c.getAssignExpr(ident.Obj)
	c.checkExpr(expr, pos)
}
