package ch2

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)

func TestAST1(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+2*3`)
	ast.Print(nil, expr)
}

func TestAST2(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+2*3`)
	ast.Print(nil, Eval(expr))
}

func Eval(exp ast.Expr) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	}
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) float64 {
	switch exp.Op {
	case token.ADD:
		return Eval(exp.X) + Eval(exp.Y)
	case token.MUL:
		return Eval(exp.X) * Eval(exp.Y)
	}
	return 0
}

func TestAST3(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+2*3+x`)
	fmt.Println(EvalV2(expr, map[string]float64{"x": 100}))
}

func EvalV2(exp ast.Expr, vars map[string]float64) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExprV2(exp, vars)
	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	case *ast.Ident:
		return vars[exp.Name]
	}
	return 0
}

func EvalBinaryExprV2(exp *ast.BinaryExpr, vars map[string]float64) float64 {
	switch exp.Op {
	case token.ADD:
		return EvalV2(exp.X, vars) + EvalV2(exp.Y, vars)
	case token.MUL:
		return EvalV2(exp.X, vars) * EvalV2(exp.Y, vars)
	}
	return 0
}
