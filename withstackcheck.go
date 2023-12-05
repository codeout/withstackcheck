package withstackcheck

import (
	"go/ast"

	"github.com/k0kubun/pp"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"

	"github.com/codeout/withstackcheck/inspector"
)

const doc = "withstackcheck checks that errors from external packages are wrapped with stacktrace, and that errors from internal packages are not wrapped with stacktrace."

var Analyzer = &analysis.Analyzer{
	Name: "withstackcheck",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	checker := inspector.New(
		pass,
		"error without stacktrace returned from external package",
		"error with stacktrace returned from internal package",
	)

	// find func()
	checker.PreorderedFuncDecl(func(fnNode ast.Node) {
		if _, ok := fnNode.(*ast.FuncDecl); !ok {
			return
		}

		pp.Println(fnNode, "<<< whole func")
		checker.CheckErrorReturns(fnNode)
	})

	return nil, nil
}
