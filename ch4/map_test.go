package main

import (
	"fmt"
	"testing"
)

func k(list []string) string { return fmt.Sprintf("%q", list) }

var m = make(map[string]int)

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func TestMap(t *testing.T) {
	var name = "Bob"
	var ages = make(map[string]int)
	ages[name] = 30

	fmt.Println(ages[name])

	list := []string{"abc", "123"}
	Add(list)
	fmt.Println(k(list))
}
