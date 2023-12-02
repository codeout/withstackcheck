package withstackcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "withstackcheck checks that errors from external packages are wrapped with stacktrace, and that errors from internal packages are not wrapped with stacktrace."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "withstackcheck",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspctr := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspctr.Preorder(nodeFilter, func(n ast.Node) {
		if n, ok := n.(*ast.Ident); ok {
			if n.Name == "gopher" {
				pass.Reportf(n.Pos(), "identifier is gopher")
			}
		}
	})

	return nil, nil
}
