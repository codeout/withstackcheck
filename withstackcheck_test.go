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
}
