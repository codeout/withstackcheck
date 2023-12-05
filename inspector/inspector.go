package inspector

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var funcDeclFilter = []ast.Node{
	(*ast.FuncDecl)(nil),
}

type WithStackChecker struct {
	pass              *analysis.Pass
	inspector         *inspector.Inspector
	withoutStackError string
	withStackError    string
}

func (c *WithStackChecker) PreorderedFuncDecl(f func(ast.Node)) {
	c.inspector.Preorder(funcDeclFilter, f)
}

// CheckErrorReturns returns all return statements that return an error.
func (c *WithStackChecker) CheckErrorReturns(fnNode ast.Node) {
	ast.Inspect(fnNode, func(node ast.Node) bool {
		ret, ok := node.(*ast.ReturnStmt)
		if !ok {
			return true
		}

		for _, expr := range ret.Results {
			if !c.isError(expr) {
				continue
			}

			c.checkExpr(expr, expr.Pos())
		}

		return false
	})
}

// checkExpr checks the generic expression and report.
func (c *WithStackChecker) checkExpr(expr ast.Expr, pos token.Pos) {
	switch expr := expr.(type) {
	case *ast.Ident:
		c.checkIdent(expr, pos)
	case *ast.CallExpr:
		c.checkCallExpr(expr, pos)
	default:
		panic(fmt.Sprintf("Unimplemented type: %T", expr))
	}
}

func (c *WithStackChecker) isError(expr ast.Expr) bool {
	typ := c.pass.TypesInfo.TypeOf(expr)

	if typ == nil {
		return false
	}

	return typ.String() == "error"
}

func New(pass *analysis.Pass, withoutStackError string, withStackError string) *WithStackChecker {
	return &WithStackChecker{
		pass:              pass,
		inspector:         pass.ResultOf[inspect.Analyzer].(*inspector.Inspector),
		withoutStackError: withoutStackError,
		withStackError:    withStackError,
	}
}
