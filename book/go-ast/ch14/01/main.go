package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/constant"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"

	"golang.org/x/tools/go/ssa"
)

const src = `
package main

func main() {
	println("Hello，凹语言！")
	println("The answer is:", 42)
}
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", src, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}

	conf := types.Config{Importer: nil}
	pkg, err := conf.Check("test.go", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatal(err)
	}

	var ssaProg = ssa.NewProgram(fset, ssa.SanityCheckFunctions)
	var ssaPkg = ssaProg.CreatePackage(pkg, []*ast.File{f}, info, true)

	ssaPkg.Build()
	ssaPkg.Func("main").WriteTo(os.Stdout)

	runFunc(ssaPkg.Func("main"))
}

func runFunc(fn *ssa.Function) {
	fmt.Println("--- runFunc begin ---")
	defer fmt.Println("--- runFunc end ---")

	if len(fn.Blocks) > 0 {
		for blk := fn.Blocks[0]; blk != nil; {
			blk = runFuncBlock(fn, blk)
		}
	}
}

func runFuncBlock(fn *ssa.Function, block *ssa.BasicBlock) (nextBlock *ssa.BasicBlock) {
	for _, ins := range block.Instrs {
		switch ins := ins.(type) {
		case *ssa.Call:
			doCall(ins)
		case *ssa.Return:
			doReturn(ins)
		default:
			doUnknown(ins)
		}
	}
	return nil
}

func doCall(ins *ssa.Call) {
	switch {
	case ins.Call.Method == nil:
		switch callFn := ins.Call.Value.(type) {
		case *ssa.Builtin:
			callBuiltin(callFn, ins.Call.Args...)
		default:
		}
	default:
	}
}

func doReturn(ins *ssa.Return) {
	return // ins.Results[...]
}

func doUnknown(ins ssa.Instruction) {
	// 其它指令
	// 循环和分支结构需要处理 phi 指令
	// 目前的例子只有单个 block
}

func callBuiltin(fn *ssa.Builtin, arg ...ssa.Value) {
	switch fn.Name() {
	case "println":
		var buf bytes.Buffer
		for i := 0; i < len(arg); i++ {
			if i > 0 {
				buf.WriteRune(' ')
			}
			switch arg := arg[i].(type) {
			case *ssa.Const:
				if t, ok := arg.Type().Underlying().(*types.Basic); ok {
					switch t.Kind() {
					case types.Int, types.UntypedInt:
						fmt.Fprintf(&buf, "%d", int(arg.Int64()))
					case types.String:
						fmt.Fprintf(&buf, "%s", constant.StringVal(arg.Value))
					default:
					}
				}
			default:
			}
		}
		buf.WriteRune('\n')
		os.Stdout.Write(buf.Bytes())
	default:
	}
}
