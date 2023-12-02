package main

import (
	"github.com/codeout/withstackcheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(withstackcheck.Analyzer) }
