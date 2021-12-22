package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	parseTypeDecl()
}

func parseTypeDecl() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", typeDeclSrc, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				fmt.Printf("%T\n", spec)
			}
		}
	}
}

const typeDeclSrc = `package foo
type MyInt1 int
type MyInt2 = int
`
