package nils

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	i int
}

func NilOrNot(v interface{}) bool {
	if v != nil {
		fmt.Println(v)
		return false
	}
	return true
}

func TestNil(t *testing.T) {
	var s *TestStruct
	fmt.Println(s == nil)
	fmt.Println(NilOrNot(s))
}
