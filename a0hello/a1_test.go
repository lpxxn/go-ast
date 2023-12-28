package a0hello

import (
	"go/ast"
	"go/parser"
	"testing"
)

func TestAST1(t *testing.T) {
	expr, _ := parser.ParseExpr(`1`)
	ast.Print(nil, expr)

	expr, _ = parser.ParseExpr(`"abc"`)
	ast.Print(nil, expr)
}

func TestExp1(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+2*3`)
	ast.Print(nil, expr)
}
