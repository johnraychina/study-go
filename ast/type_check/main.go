package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", typeSrc, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	pkg, err := new(types.Config).Check("hello.go", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	_ = pkg
}

const typeSrc = `package pkg

func hello() {
	var _ = "a" + 1
}
`
