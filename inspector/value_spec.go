package inspector

import (
	"go/ast"
)

// checkValueSpec checks the value spec and report.
func (c *WithStackChecker) checkValueSpec(spec *ast.ValueSpec) {
	// need to find var assignments when "var err error" is found
	e := c.findAssignExprInFunction(spec)
	c.checkExpr(e)
}
