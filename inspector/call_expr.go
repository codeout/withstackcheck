package inspector

import (
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
		case *ast.SelectorExpr:
			if !c.isExternalPackage(expr) {
				c.pass.Reportf(c.pos, c.withStackError)
			}
		case *ast.CallExpr:
			if !c.isExternalPackage(expr) {
				c.pass.Reportf(c.pos, c.withStackError)
			}
		default:
			c.panicf("Unimplemented type: %T at %s", expr, c.pass.Fset.Position(c.pos))
		}

		return
	}

	if c.isExternalPackage(callExpr) {
		c.pass.Reportf(c.pos, c.withoutStackError)
	}
}

func (c *WithStackChecker) enterWithStack() {
	c.inWithStack = true
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
