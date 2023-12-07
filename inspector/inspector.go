package inspector

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type WithStackChecker struct {
	pass              *analysis.Pass
	inspector         *inspector.Inspector
	withoutStackError string
	withStackError    string

	// current context
	funcNode    *ast.FuncDecl
	pos         token.Pos // original position of the return statement
	inWithStack bool
}

func (c *WithStackChecker) PreorderedFuncDecl(f func(ast.Node)) {
	c.inspector.Preorder(
		[]ast.Node{(*ast.FuncDecl)(nil)},
		f,
	)
}

// CheckErrorReturns returns all return statements that return an error.
func (c *WithStackChecker) CheckErrorReturns(fnNode *ast.FuncDecl) {
	ast.Inspect(fnNode, func(node ast.Node) bool {
		ret, ok := node.(*ast.ReturnStmt)
		if !ok {
			return true
		}

		for _, expr := range ret.Results {
			if !c.isError(expr) {
				continue
			}

			c.setContext(fnNode, ret.Pos()) // set current context
			c.checkExpr(expr)
		}

		return false
	})
}

// checkExpr checks the generic expression and report.
func (c *WithStackChecker) checkExpr(expr ast.Expr) {
	switch expr := expr.(type) {
	case *ast.Ident:
		c.checkIdent(expr)
	case *ast.CallExpr:
		c.checkCallExpr(expr)
	default:
		panic(fmt.Sprintf("Unimplemented type: %T", expr))
	}
}

func (c *WithStackChecker) setContext(fnNode *ast.FuncDecl, pos token.Pos) {
	c.funcNode = fnNode
	c.pos = pos
	c.inWithStack = false
}

func (c *WithStackChecker) enterWithStack() {
	c.inWithStack = true
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
