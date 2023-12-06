package inspector

import (
	"go/ast"
)

// checkIdent checks the identifier and report.
func (c *WithStackChecker) checkIdent(ident *ast.Ident) {
	// need to find var assignments when "var err error" is found
	if spec, ok := ident.Obj.Decl.(*ast.ValueSpec); ok {
		c.checkValueSpec(spec)
		return
	}

	expr := c.getAssignExprInObject(ident.Obj)
	c.checkExpr(expr)
}
