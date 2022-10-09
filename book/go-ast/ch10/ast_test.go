package ch10

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAST1(t *testing.T) {
	var src = `
package pkgname

func main() {}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)

	src = `
package pkgname

func main() {
	{}
	{}
}
`
	f, err = parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

func TestAST2(t *testing.T) {
	var src = `
package pkgname

func main() {
	42
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

func TestAST3(t *testing.T) {
	var src = `
package pkgname

func main() {
	return 42, err
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

func TestAST4(t *testing.T) {
	var src = `
package pkgname

func main() {
	var a int
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

func TestAST5(t *testing.T) {
	var src = `
package pkgname

func main() {
	a, b := 1, 2
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

func TestAST6(t *testing.T) {
	var src = `
package pkgname

func main() {
	if true {} else {}
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

func TestAST7(t *testing.T) {
	var src = `
package pkgname

func main() {
	for {}
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)

	src = `
package pkgname

func main() {
	for true {}
}
`
	f, err = parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)

	src = `
package pkgname

func main() {
	for x; y; z {}
}
`
	f, err = parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)

	src = `
package pkgname

func main() {
	for range ch {}
}
`
	f, err = parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

func TestAST8(t *testing.T) {
	var src = `
package pkgname

func main() {
	x.(int)
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

func TestAST9(t *testing.T) {
	var src = `
package pkgname

func main() {
	go hello("hi world")
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		t.Error(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}
