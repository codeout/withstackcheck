package inspector

import (
	"fmt"
	"go/ast"
	"go/token"
)

// checkCallExpr checks the call expression and report.
// It requires pos which looks unnecessary, but it's needed to report the original position.
// checkIndent() strips expressions and also the original position for example.
func (c *WithStackChecker) checkCallExpr(callExpr *ast.CallExpr, pos token.Pos) {
	if c.isWithStack(callExpr) {
		// don't call other checkXXX() inside as it's an argument of errors.WithStack()
		switch expr := callExpr.Args[0].(type) {
		case *ast.Ident:
			e := c.getAssignExpr(expr.Obj)
			if !c.isExternalPackage(e) {
				c.pass.Reportf(pos, c.withStackError)
			}
		default:
			panic(fmt.Sprintf("Unimplemented type: %T", expr))
		}

		return
	}

	if c.isExternalPackage(callExpr) {
		c.pass.Reportf(pos, c.withoutStackError)
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

// isExternalPackage checks if the expression is external package call.
func (c *WithStackChecker) isExternalPackage(expr ast.Expr) bool {
	callExpr, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}
	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	return c.pass.TypesInfo.ObjectOf(selExpr.Sel).Pkg().Path() != c.pass.Pkg.Path()
}
