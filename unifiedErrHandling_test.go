package unifiedErrHandling_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/mercy34mercy/unifiedErrHandling"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
// func TestAnalyzer(t *testing.T) {
// 	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
// 	analysistest.Run(t, testdata, unifiedErrHandling.Analyzer, "a")
// }

func TestAnalyzer2(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, unifiedErrHandling.Analyzer, "b")
}

func TestAnalyzer3(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, unifiedErrHandling.Analyzer, "c")
}
