package main

import (
	"github.com/mercy34mercy/unifiedErrHandling"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unifiedErrHandling.Analyzer) }
