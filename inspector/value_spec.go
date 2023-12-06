package inspector

import (
	"go/ast"
)

// checkValueSpec checks the value spec and report.
func (c *WithStackChecker) checkValueSpec(spec *ast.ValueSpec) {
	e := c.getAssignExprFromValueSpec(spec)
	c.checkExpr(e)
}
