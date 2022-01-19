package main

import (
	"errors"
	"fmt"
)

func main() {
	// 采用short variable声明，范围是block，if内部和外部是不同的
	v1, err := f1()
	if v1 > 0 {
		_, err := f2()
		if err != nil {
		}
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
}

func f2() (int, error) {
	return 2, errors.New("biz error")
}
func f1() (int, error) {
	return 1, nil
}
