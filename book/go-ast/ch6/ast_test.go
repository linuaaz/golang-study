package ch6

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAST1(t *testing.T) {
	var src = `
package hello

func (p *xType) Hello(arg1, arg2 int) (bool, error) { return false, nil }
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			ast.Print(nil, fn)
		}
	}
}

func TestAST2(t *testing.T) {
	var src = `
package hello

func Hello1(s0, s1 string, s2 string) {}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			ast.Print(nil, fn)
		}
	}
}
