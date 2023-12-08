package inspector

import (
	"go/ast"
)

// checkSelectorExpr checks the selector expression and report.
func (c *WithStackChecker) checkSelectorExpr(selExpr *ast.SelectorExpr) {
	// this is always called out of "errors.WithStack()"
	if c.isExternalPackage(selExpr) {
		c.pass.Reportf(c.pos, c.withoutStackError)
		return
	}
}
