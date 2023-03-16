package unifiedErrHandling

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "test is ..."

var (
	abbreviatedCnt = 0
	SeparatedCnt   = 0
)

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "test",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.IfStmt:
			if _, ok := n.Cond.(*ast.BinaryExpr); ok {
				lhs := n.Cond.(*ast.BinaryExpr).X
				rhs := n.Cond.(*ast.BinaryExpr).Y

				if isErrVar(pass.TypesInfo, lhs) && isNil(rhs) {
					if n.Init != nil {
						abbreviatedCnt++
					} else {
						SeparatedCnt++
					}
				}
			}
		}
	})

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.IfStmt:
			if _, ok := n.Cond.(*ast.BinaryExpr); ok {
				lhs := n.Cond.(*ast.BinaryExpr).X
				rhs := n.Cond.(*ast.BinaryExpr).Y

				if isErrVar(pass.TypesInfo, lhs) && isNil(rhs) {
					if n.Init != nil {
						if abbreviatedCnt < SeparatedCnt {
							pass.Reportf(n.Pos(), "abbreviated notation")
						}
					} else {
						if SeparatedCnt < abbreviatedCnt {
							pass.Reportf(n.Pos(), "Separated notation")
						}
					}
				}
			}
		}
	})

	return nil, nil
}

func isErrVar(info *types.Info, expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		if types.Identical(info.ObjectOf(ident).Type(), types.Universe.Lookup("error").Type()) {
			return true
		}
	}
	return false
}

func isNil(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok && ident.Obj == nil {
		return true
	}
	return false
}
