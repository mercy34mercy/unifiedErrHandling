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

	//何を使いたいのか選択、今回は、*ast.IFStmt これはIf文の情報がとれるやつ
	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
	}

	//Preordreで探してきて、今回はIFStmtだけだが、型がanyなのでcaseを使ってキャストしている
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.IfStmt:
			// IfStmtの構文木に *ast.BinaryExprがありこれを使うと、比較されている二つの変数がとれる
			if _, ok := n.Cond.(*ast.BinaryExpr); ok {
				//左側
				lhs := n.Cond.(*ast.BinaryExpr).X
				//右側
				rhs := n.Cond.(*ast.BinaryExpr).Y
				
				//比較対象がerror型とnilかどうかを判別
				if isErrVar(pass.TypesInfo, lhs) && isNil(rhs){
					//n.Initでif文の中で変数定義がされているか確認
					if n.Init != nil {
						abbreviatedCnt++
					} else {
						SeparatedCnt++
					}
				}else if(isErrVar(pass.TypesInfo, rhs) && isNil(lhs)){
					if n.Init != nil {
						abbreviatedCnt++
					} else {
						SeparatedCnt++
					}
				}
			}
		}
	})

	//数を数えてどっちにエラーを出すのか決めたかったので、二回同じ処理を回す
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
							pass.Reportf(n.Pos(), "separated notation")
						}
					}
				}
			}
		}
	})

	return nil, nil
}

//その値がerror型か確かめる プログラムの型の情報とast木の中の調べたい変数を渡す
func isErrVar(info *types.Info, expr ast.Expr) bool {
	//変数はだいたいこの *ast.Expr型なのでキャスト
	if ident, ok := expr.(*ast.Ident); ok {
		//types.Identicalで型が同じか確かめることができる、引数1が今回調べたい変数の型を調べて、引数2でerror型を指定している
		if types.Identical(info.ObjectOf(ident).Type(), types.Universe.Lookup("error").Type()) {
			return true
		}
	}
	return false
}

//その値がnilかどうか確かめる
func isNil(expr ast.Expr) bool {
	//処理としては、errror型の時と似ているが、
	if ident, ok := expr.(*ast.Ident); ok && ident.Obj == nil {
		return true
	}
	return false
}
