package ch4

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

const src = `
package pkgname

import ("a"; "b")
type SomeType int
const PI = 3.14
var Length = 1

func main() {}
`

func TestAST1(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}
	ast.Print(nil, f)
}

func TestAST2(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("package:%s\n", f.Name)

	for _, s := range f.Imports {
		t.Logf("import:%s\n", s.Path.Value)
	}

	for _, decl := range f.Decls {
		t.Logf("decl: %T\n", decl)
	}

	for _, v := range f.Decls {
		if s, ok := v.(*ast.GenDecl); ok && s.Tok == token.IMPORT {
			for _, v := range s.Specs {
				t.Logf("import:%s\n", v.(*ast.ImportSpec).Path.Value)
			}
		}
	}
}

type myNodeVisitor struct{}

func (p *myNodeVisitor) Visit(n ast.Node) (w ast.Visitor) {
	if x, ok := n.(*ast.Ident); ok {
		fmt.Println("myNodeVisitor.Visit:", x.Name)
	}
	return p
}

func TestAST3(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Walk(new(myNodeVisitor), f)
}

func TestAST4(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if x, ok := n.(*ast.Ident); ok {
			fmt.Println("ast.Inspect:", x.Name)
		}
		return true
	})
}
