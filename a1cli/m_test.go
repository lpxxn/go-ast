package main_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAST1(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, parser.AllErrors)
	if err != nil {
		t.Log(err)
	}
	t.Logf("package %s", f.Name)
	for _, s := range f.Imports {
		t.Log("import: ", s.Path.Value)
	}
	for _, decl := range f.Decls {
		t.Logf("decl: %T detail:  %#v", decl, decl)
	}

	/*
		但是结构体中最重要的其实是File.Decls成员，它包含了当前文件全部的包级声明信息（包含导入信息）。
		即使没有File.Imports成员，我们也可以从File.Decls声明列表中获取全部导入包的信息。
	*/
	for _, v := range f.Decls {
		if s, ok := v.(*ast.GenDecl); ok && s.Tok == token.IMPORT {
			for _, v := range s.Specs {
				t.Log("import: ", v.(*ast.ImportSpec).Path.Value)
			}
		}
	}
}

const src = `package pkgname

import ("a"; "b")
type SomeType int
const PI = 3.14
var Length = 1

func main() {}
`
