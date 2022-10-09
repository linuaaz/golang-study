package ch2

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAST1(t *testing.T) {
	var lit9527 = &ast.BasicLit{
		Kind:  token.INT,
		Value: "9527",
	}
	ast.Print(nil, lit9527)
}

func TestAST2(t *testing.T) {
	expr, _ := parser.ParseExpr(`9527`)
	ast.Print(nil, expr)

	expr, _ = parser.ParseExpr(`"9527"`)
	ast.Print(nil, expr)
}

func TestAST3(t *testing.T) {
	ast.Print(nil, ast.NewIdent(`x`))

	expr, _ := parser.ParseExpr(`x`)
	ast.Print(nil, expr)
}
