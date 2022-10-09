package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

type Program struct {
	fs   map[string]string
	ast  map[string]*ast.File
	pkgs map[string]*types.Package
	fset *token.FileSet
}

func NewProgram(fs map[string]string) *Program {
	return &Program{
		fs:   fs,
		ast:  make(map[string]*ast.File),
		pkgs: make(map[string]*types.Package),
		fset: token.NewFileSet(),
	}
}

func (p *Program) LoadPackage(path string) (pkg *types.Package, f *ast.File, err error) {
	if pkg, ok := p.pkgs[path]; ok {
		return pkg, p.ast[path], nil
	}

	f, err = parser.ParseFile(p.fset, path, p.fs[path], parser.AllErrors)
	if err != nil {
		return nil, nil, err
	}

	conf := types.Config{Importer: p}
	pkg, err = conf.Check(path, p.fset, []*ast.File{f}, nil)
	if err != nil {
		return nil, nil, err
	}

	p.ast[path] = f
	p.pkgs[path] = pkg
	return pkg, f, nil
}

func (p *Program) Import(path string) (*types.Package, error) {
	if pkg, ok := p.pkgs[path]; ok {
		return pkg, nil
	}
	// V1
	//return nil, fmt.Errorf("not found: %s", path)
	// V2
	pkg, _, err := p.LoadPackage(path)
	return pkg, err
}

func main() {
	prog := NewProgram(map[string]string{
		"hello": `
            package main
            import "math"
            func main() { var _ = 2 * math.Pi }
        `,
		"math": `
            package math
            const Pi = 3.1415926
        `,
	})

	_, f, err := prog.LoadPackage("math")
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(nil, f)

	_, f, err = prog.LoadPackage("hello")
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(nil, f)
}
