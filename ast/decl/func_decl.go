package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	parseFuncDecl()
}

func parseFuncDecl() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", funcDecl, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			ast.Print(nil, fn)
		}
	}
}

const funcDecl = `package foo
func (p *xType) Hello(arg1, arg2 int) (bool, error) {
	return true, nil
}
`
