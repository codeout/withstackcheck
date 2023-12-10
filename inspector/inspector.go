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
	funcNode     *ast.FuncDecl
	namedReturns []*ast.Ident
	pos          token.Pos // original position of the return statement
	inWithStack  bool
}

func (c *WithStackChecker) PreorderedFuncDecl(f func(ast.Node)) {
	c.inspector.Preorder(
		[]ast.Node{(*ast.FuncDecl)(nil)},
		f,
	)
}

// CheckErrorReturns returns all return statements that return an error.
func (c *WithStackChecker) CheckErrorReturns(fnNode *ast.FuncDecl) {
	// Find named returns. If it doesn't exist, set empty slice.
	c.setNamedReturns(fnNode.Type.Results.List)

	ast.Inspect(fnNode, func(node ast.Node) bool {
		ret, ok := node.(*ast.ReturnStmt)
		if !ok {
			return true
		}

		for _, expr := range ret.Results {
			c.check(fnNode, expr)
		}

		// if black "return" is found and there is any named return, handle the return as named
		if len(ret.Results) == 0 {
			for _, namedRet := range c.namedReturns {
				namedRet.NamePos = ret.Pos() // use position of the "return"
				c.check(fnNode, namedRet)
			}
		}

		return false
	})
}

func (c *WithStackChecker) check(fnNode *ast.FuncDecl, expr ast.Expr) {
	if _, ok := expr.(*ast.CallExpr); !ok && !c.isError(expr) {
		return
	}

	c.setContext(fnNode, expr.Pos()) // set current context
	c.checkExpr(expr)
}

// checkExpr checks the generic expression and report.
func (c *WithStackChecker) checkExpr(expr ast.Expr) {
	switch expr := expr.(type) {
	case nil:
		return // mark passed
	case *ast.Ident:
		c.checkIdent(expr)
	case *ast.CallExpr:
		c.checkCallExpr(expr)
	case *ast.SelectorExpr:
		c.checkSelectorExpr(expr)
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

func (c *WithStackChecker) setNamedReturns(fields []*ast.Field) {
	var namedReturns []*ast.Ident
	for _, field := range fields {
		namedReturns = append(namedReturns, field.Names...)
	}

	c.namedReturns = namedReturns
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
