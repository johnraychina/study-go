package main

import (
	"golang.org/x/exp/constraints"
	"reflect"
)

// 泛型介绍： https://go.dev/blog/intro-generics
// 泛型功能加了3个东西：
// 1. Type Parameters for functions and types.
//  函数和类型(主要是struct)加了类型参数
// 2. Defining interface types as sets of types, including types that don't have methods.
//  把接口类型作为类型集合（即使有的类型没有方法）
// 3. Type inference, which permits omitting type arguments in many cases when calling a function.
// 增加了类型推断，允许在调用函数时，多数情况下可以省略类型参数。

func main() {

	// Type Parameters for functions
	fmin := GMin[float64]
	println("fmin:", fmin(2.71828, 3.14159))

	// Type Inference 类型推断
	// function argument type inference: 调用函数时，将m = GMin[float64](a, b) 简化为 m = GMin(a, b)
	x := GMin(2, 3)
	println("min:", x)

	// Constraint type inference: 判断Point遵从 [S ~[]E, E constraints.Integer] 的约束。
	p := Point{1, 2}
	sp := Scale(p, 2)
	println("scaled point:", sp[0], ",", sp[1])

}

// when to use generics 泛型使用时机: 先写
// https://go.dev/blog/when-generics

func ListStringKeys(m map[string]any) []string {
	s := make([]string, len(m))
	for k, _ := range m {
		s = append(s, k)
	}
	return s
}

// ListKeys collect keys of a map into a slice
func ListKeys[Key comparable](m map[Key]any) []Key {
	s := make([]Key, len(m))
	for k, _ := range m {
		s = append(s, k)
	}
	return s
}

// Type Sets 类型集：
// 对于类型参数为 T，满足约束 C（type set） 的泛型函数内，
// 泛型对象x可以执行的方法只能是对应type set中所有类型method的交集
// 例如 [S interface{~[]E}, E interface{}] 表示类型参数S 必须是一个slice类型，slice的元素可以是任意类型。
// 另外，由于这是比较场景的场景，go支持去掉包装的 interface{}，简化为： [S ~[]E, E interface{}]
// 进一步，简化为[S ~[]E, E any]

type Point []int32

func (p Point) String() string {
	// Details not important.
	return ""
}

// Scale returns a copy of s with each element multiplied by c.
func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

//func Scale[E constraints.Integer](s []E, c E) []E {
//	r := make([]E, len(s))
//	for i, v := range s {
//		r[i] = v * c
//	}
//	return r
//}

// 两种视角来看待
// method sets视角：
//	译器做合法校验：求类型参数T 的约束，例如constraints.Ordered 与实际类型的方法做判断，看是否被包含，例如float64。
// type set视角：
//	编译器做合法校验：直接使用类型定义做比较，非常简单直观 type Ordered interface{ Integer | Float | ~string}。
// ~string 波浪号代表任何底层是string的类型都算

// GMin function type parameter T
// Type Set: interface{ Integer | Float | ~string }
func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	if reflect.DeepEqual(t.value, x) {
		return t
	}
	if t.left != nil {
		if res := t.left.Lookup(x); res != nil {
			return res
		}
	}
	if t.right != nil {
		if res := t.right.Lookup(x); res != nil {
			return res
		}
	}
	return nil
}
