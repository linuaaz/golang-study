package ch7

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAST1(t *testing.T) {
	var src = `
package foo

type Int1 int
type Int2 pkg.int
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	for _, decl := range f.Decls {
		ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
	}
}

func TestAST2(t *testing.T) {
	var src = `
package foo

type IntPtr *int
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	for _, decl := range f.Decls {
		ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
	}
}

func TestAST3(t *testing.T) {
	var src = `
package foo

type IntPtrPtr **int
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	for _, decl := range f.Decls {
		ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
	}
}

func TestAST4(t *testing.T) {
	var src = `
package foo

type IntArray [1]int
type IntArray2 [...]int
type IntArrayArray [1][2]int
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	for _, decl := range f.Decls {
		ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
	}
}

func TestAST5(t *testing.T) {
	var src = `
package foo

type IntSlice []int
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	for _, decl := range f.Decls {
		ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
	}
}

func TestAST6(t *testing.T) {
	var src = `
package foo

type MyStruct struct {
    a, b int "int value"
    string
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f)
}

func TestAST7(t *testing.T) {
	var src = `
package foo

type IntStringMap map[int]string
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f)
}

func TestAST8(t *testing.T) {
	var src = `
package foo

type IntChan chan int
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f)
}

func TestAST9(t *testing.T) {
	var src = `
package foo

type FuncType func(a, b int) bool]
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f)
}

func TestAST10(t *testing.T) {
	var src = `
package foo

type IntReader interface {
    Read() int
}

type IntReader struct {
    Read func() int
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	for _, decl := range f.Decls {
		ast.Print(nil, decl.(*ast.GenDecl).Specs[0])
	}
}
