package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := Status(10)
	fmt.Println(fmt.Sprintf("%s", s))

	s2 := Status(20)
	fmt.Println(s2)

	s3 := Status(30)
	fmt.Println(s3)
}

type Status int32

const (
	Pending Status = 0
	Success Status = 10
	Fail    Status = 20
)

func (s Status) String() string {
	switch s {
	case Pending:
		return "0-处理中"
	case Success:
		return "10-成功"
	case Fail:
		return "20-失败"
	default:
		return strconv.Itoa(int(s))
	}
}
