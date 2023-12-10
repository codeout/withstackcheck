package withstackcheck_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/codeout/withstackcheck"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/assign_var")
	analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/reassign_var")
	analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/named_return")
	analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/other_package")
	analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/interface")
	analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/func")
	analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/anonymous_func")
	analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/anonymous_func_with_mixed_var_scope")
	analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/type_assert")
	// TODO: implement check for *ast.indexExpr
	// analysistest.Run(t, testdata, withstackcheck.Analyzer, "withstackcheck/index")
}
