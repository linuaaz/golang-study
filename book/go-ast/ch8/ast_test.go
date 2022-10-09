package ch8

import (
	"go/ast"
	"go/parser"
	"testing"
)

func TestAST1(t *testing.T) {
	expr, _ := parser.ParseExpr(`func(){}`)
	ast.Print(nil, expr)
}

func TestAST2(t *testing.T) {
	expr, _ := parser.ParseExpr(`
[1]int{1}
[...]int{100:1,200:2}
[]int{1,2,3}
[]int{100:1,200:2}
struct {X int}{1}
struct {X int}{X:1}
map[int]int{1:1, 2:2}
`)
	ast.Print(nil, expr)
}

func TestAST3(t *testing.T) {
	expr, _ := parser.ParseExpr(`[...]int{1,2:3}`)
	ast.Print(nil, expr)
}

func TestAST4(t *testing.T) {
	expr, _ := parser.ParseExpr(`struct{X int}{X:1}`)
	ast.Print(nil, expr)

	expr, _ = parser.ParseExpr(`struct{X int}{1}`)
	ast.Print(nil, expr)
}

func TestAST5(t *testing.T) {
	expr, _ := parser.ParseExpr(`map[int]int{1:2}`)
	ast.Print(nil, expr)
}
