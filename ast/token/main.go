package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"strconv"
)

func main() {
	//parseExpr()
	//ident()
	//expression()
	//expression2()
	//fileOrganization()
	//walk()
	inspect()
}

func inspect() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	ast.Inspect(f, func(node ast.Node) bool {
		if x, ok := node.(*ast.Ident); ok {
			fmt.Println("myNodeVisitor.Visit:", x.Name)
		}
		return true
	})
}

type myNodeVisitor struct{}

func (m myNodeVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if x, ok := node.(*ast.Ident); ok {
		fmt.Println("myNodeVisitor.Visit:", x.Name)
	}
	return m
}

func walk() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	ast.Walk(new(myNodeVisitor), f)
}

const src = `package pkgname

import ("a"; "b")
type SomeType int
const PI = 3.14
var Length = 1

func main() {}
`

func fileOrganization() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("package:", f.Name)
	for _, s := range f.Imports {
		fmt.Println("import", s.Path.Value)
	}

	for _, decl := range f.Decls {
		fmt.Printf("decl: %T\n", decl)
	}

	for _, v := range f.Decls {
		if s, ok := v.(*ast.GenDecl); ok && s.Tok == token.IMPORT {
			for _, v := range s.Specs {
				fmt.Println("import:", v.(*ast.ImportSpec).Path.Value)
			}
		}
	}
}

func expression2() {
	expr, _ := parser.ParseExpr(`1+2*3+x`)
	fmt.Println(Eval(expr, map[string]float64{
		"x": 100,
	}))
}

func Eval(exp ast.Expr, vars map[string]float64) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp, vars)
	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	case *ast.Ident:
		return vars[exp.Name]
	}
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr, vars map[string]float64) float64 {
	switch exp.Op {
	case token.ADD:
		return Eval(exp.X, vars) + Eval(exp.Y, vars)
	case token.MUL:
		return Eval(exp.X, vars) * Eval(exp.Y, vars)
	}
	return 0
}

func expression() {
	expr, _ := parser.ParseExpr(`1+2*3`)
	ast.Print(nil, expr)
}
func ident() {
	ast.Print(nil, ast.NewIdent(`x`))
	expr, _ := parser.ParseExpr(`x := 3`)
	ast.Print(nil, expr)
}

func parseExpr() {
	expr, _ := parser.ParseExpr(`9527`)
	ast.Print(nil, expr)
}

func scan() {
	var src = []byte(`println("你好，世界")`)

	var fset = token.NewFileSet()
	var file = fset.AddFile("hello.go", fset.Base(), len(src))

	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
func goAst() {
	src := `
package main
func main() {
	print(1+1)
}
`
	// Create the AST by parsing src.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}

func jump(nums []int) int {

	// 最优化问题求解
	// 目标: min step
	// 定义: dp[i] 从开始位置到达位置i最小跳跃次数，dp[n-1] 即为所求解
	// 递推公式：dp[i] = min{ if nums[i-1]>=1: dp[i-1]+1, if num[i-2]>=2: dp[i-2]+1, ..., if num[i-x]>=x: dp[i-x]+1  }
	// sample: nums = [2,3,1,1,4]

	n := len(nums)
	dp := make([]int, n, n)
	// 先初始化为极大值
	for i := 1; i < n; i++ {
		dp[i] = 1000
	}

	for i := 1; i < n; i++ {
		// 得到  dp[i]
		for x := 1; x <= i; x++ {
			if (nums[i-x] >= x) && (dp[i-x]+1 < dp[i]) {
				dp[i] = dp[i-x] + 1
			}
		}
	}

	return dp[n-1]
}
