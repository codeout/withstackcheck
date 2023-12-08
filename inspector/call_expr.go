package inspector

import (
	"fmt"
	"go/ast"
)

// checkCallExpr checks the call expression and report.
func (c *WithStackChecker) checkCallExpr(callExpr *ast.CallExpr) {
	if c.inWithStack || c.isWithStack(callExpr) {
		c.enterWithStack()

		switch expr := callExpr.Args[0].(type) {
		case *ast.Ident:
			e := c.findAssignExprInFunction(expr.Obj.Decl)
			if !c.isExternalPackage(e) {
				c.pass.Reportf(c.pos, c.withStackError)
			}
		default:
			panic(fmt.Sprintf("Unimplemented type: %T", expr))
		}

		return
	}

	if c.isExternalPackage(callExpr) {
		c.pass.Reportf(c.pos, c.withoutStackError)
	}
}

// isWithStack checks if the call expression is errors.WithStack().
func (c *WithStackChecker) isWithStack(expr *ast.CallExpr) bool {
	selExpr, ok := expr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if ident, ok := selExpr.X.(*ast.Ident); !ok || ident.Name != "errors" {
		return false
	}

	return selExpr.Sel.Name == "WithStack"
}
