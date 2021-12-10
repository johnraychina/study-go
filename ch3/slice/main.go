package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {

	//test1()
	test2()
	//test3()
}

func test3() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
	numbers[3] = 1
	printSlice(numbers)
}
func test2() {
	// 和数组一样，len控制可读写的元素边界，加了个cap，用于在append时判断底层数组是否需要扩容
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6}
	slice1 := arr[1:4:5]        //{low:high:max} 最多再扩张一个元素
	slice1[2] = 33              //修改, {1,2,3}, len=3, cap = 4
	slice1 = append(slice1, 44) //cap内追加, {1,2,3,44}, len=4, cap = 4
	slice1 = append(slice1, 55) //cap扩容, {1,2,3,44, 55}, len=5, cap = 8
	fmt.Println(slice1)

	x := slice1[1:2]
	fmt.Println(x)

}

func test1() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
