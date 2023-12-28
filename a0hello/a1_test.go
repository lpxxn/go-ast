package a0hello

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
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

	expr, _ = parser.ParseExpr(`1+x`)
	ast.Print(nil, expr)
}

func TestExp2(t *testing.T) {
	expr, _ := parser.ParseExpr(`x`)
	ast.Print(nil, expr)
}

func TestExp3(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+x+y`)
	ast.Print(nil, expr)
	t.Log("result: ", eval(expr, map[string]float64{"x": 2, "y": 3}))
}

func eval(exp ast.Expr, vars map[string]float64) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp, vars)

	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	case *ast.Ident:
		return vars[exp.Name]
	}
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr, vars map[string]float64) float64 {
	switch exp.Op {
	case token.ADD:
		return eval(exp.X, vars) + eval(exp.Y, vars)
	case token.MUL:
		return eval(exp.X, vars) * eval(exp.Y, vars)
	}
	return 0
}
