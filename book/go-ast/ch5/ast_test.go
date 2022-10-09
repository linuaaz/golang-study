package ch5

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

const src = `
package foo

import "pkg-a"
import pkg_b_v2 "pkg-b"
import . "pkg-c"
import _ "pkg-d"
`

func TestAST1(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.ImportsOnly)
	if err != nil {
		t.Error(err)
		return
	}

	for _, s := range f.Imports {
		fmt.Printf("import: name = %v, path = %#v\n", s.Name, s.Path)
	}
}

const src1 = `
package foo

type MyInt1 int
type MyInt2 = int
`

func TestAST2(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src1, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	//ast.Print(nil, f)

	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				t.Logf("%T\n", spec)
			}
		}
	}
}

const src2 = `
package foo

const Pi = 3.14
const E float64 = 2.71828
`

func TestAST3(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src2, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	//ast.Print(nil, f)

	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				t.Logf("%T\n", spec)
			}
		}
	}
}

const src3 = `
package foo

var Pi = 3.14
`

func TestAST4(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src3, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				ast.Print(nil, spec)
			}
		}
	}
}

const src4 = `
package foo

const Pi = 3.14

var (
    a int
    b bool
)
`

func TestAST5(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src4, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f)
}
