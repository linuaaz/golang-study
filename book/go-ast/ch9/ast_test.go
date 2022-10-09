package ch9

import (
	"go/ast"
	"go/parser"
	"testing"
)

func TestAST1(t *testing.T) {
	expr, _ := parser.ParseExpr(`int(x)`)
	ast.Print(nil, expr)
}

func TestAST2(t *testing.T) {
	expr, _ := parser.ParseExpr(`x.y`)
	ast.Print(nil, expr)
}

func TestAST3(t *testing.T) {
	expr, _ := parser.ParseExpr(`x[y]`)
	ast.Print(nil, expr)
}

func TestAST4(t *testing.T) {
	expr, _ := parser.ParseExpr(`x[1:2:3]`)
	ast.Print(nil, expr)
}

func TestAST5(t *testing.T) {
	expr, _ := parser.ParseExpr(`x.(y)`)
	ast.Print(nil, expr)
}
