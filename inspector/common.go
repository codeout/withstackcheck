package inspector

import (
	"fmt"
	"go/ast"
)

// getAssignExprToObject scans right-hands of assignments and returns the expression that is assigned to the given object.
// Note that multiple vars might be assigned in a single assignment.
func (c *WithStackChecker) getAssignExprToObject(obj *ast.Object) ast.Expr {
	switch decl := obj.Decl.(type) {
	case *ast.AssignStmt:
		return c.getAssignExprInAssignStmt(decl, obj.Decl)
	case *ast.ValueSpec:
		return c.findAssignExprInFunction(decl)
	default:
		panic(fmt.Sprintf("Unimplemented type: %T", decl))
	}
}

// getAssignExprInAssignStmt scans right-hands of assignments and returns the expression that is assigned to the given object.
func (c *WithStackChecker) getAssignExprInAssignStmt(assign *ast.AssignStmt, obj any) ast.Expr {
	switch len(assign.Rhs) {
	case 0:
		return nil
	case 1:
		return assign.Rhs[0]
	}

	if len(assign.Lhs) != len(assign.Rhs) {
		panic("Unmatched length of lhs and rhs")
	}

	for i, expr := range assign.Lhs {
		if ident, ok := expr.(*ast.Ident); ok && ident.Obj.Decl == obj {
			return assign.Rhs[i]
		}
	}

	return nil
}

// findAssignExprInFunction scans the whole function node to find the last assignment of the given spec.
func (c *WithStackChecker) findAssignExprInFunction(spec *ast.ValueSpec) ast.Expr {
	var ret ast.Expr

	ast.Inspect(c.funcNode, func(node ast.Node) bool {
		as, ok := node.(*ast.AssignStmt)
		if !ok {
			return true
		}

		// find spec in left-hands
		for _, expr := range as.Lhs {
			ident, ok := expr.(*ast.Ident)
			if !ok || ident.Obj == nil || ident.Obj.Decl != spec ||
				as.Pos() >= c.pos {
				continue
			}

			ret = c.getAssignExprInAssignStmt(as, spec)
			return false
		}

		return false
	})

	return ret
}
