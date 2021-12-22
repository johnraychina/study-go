package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	parseImportDecl()
}

func parseImportDecl() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range f.Imports {
		fmt.Printf("import: name=%v, path=%#v \n", s.Name, s.Path.Value)
	}
}

const src = `package foo
import "pkg-a"
import alias "pkg-b"
import . "pkg-c"
import _ "pkg-d"
`
