package ch1

import (
	"go/scanner"
	"go/token"
	"testing"
)

func TestAST1(t *testing.T) {
	var src = []byte(`println("你好，世界"`)

	var fset = token.NewFileSet()
	var file = fset.AddFile("hello.go", fset.Base(), len(src))

	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		t.Logf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

}
