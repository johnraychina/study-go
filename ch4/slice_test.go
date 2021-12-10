package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	//months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "March",
	//	6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	//fmt.Printf("%T", months)
	//
	//Q2 := months[4:7]
	//summer := months[6:9]
	//fmt.Println(Q2)     // ["April" "May" "June"]
	//fmt.Println(summer) // ["June" "July" "August"]

	//fmt.Println(summer[:8]) // panic: out of range

	// string 按byte个数切片，而非按unicode rune码点切片
	//每一个UTF8字符解码，不管是显式地调用utf8.DecodeRuneInString解码或是在range循环中隐式地解码，
	//如果遇到一个错误的UTF8编码输入，将生成一个特别的Unicode字符\uFFFD，在印刷中这个符号通常是一个黑色六角或钻石形状，里面包含一个白色的问号"?"。
	//当程序遇到这样的一个字符，通常是一个危险信号，说明输入并不是一个完美没有错误的UTF8字符串。
	//greet := "hello, 张先生"
	//fmt.Println(greet[7:8])

	slice := [...]int{0, 1, 2, 3, 4}
	copy(slice[3:], slice[3+1:])
	fmt.Println(slice)
}
