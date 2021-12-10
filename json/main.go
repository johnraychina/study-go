package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//list := []*Employee{
	//	&Employee{name: "Jack", age: 22},
	//	&Employee{name: "Patrick", age: 22},
	//}
	var list = []Employee{
		{Name: "Jack", Age: 22},
		{Name: "Patrick", Age: 22},
	}
	bytes, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}

	list2 := make([]*Employee, 0)
	err = json.Unmarshal(bytes, &list2)
	if err != nil {
		fmt.Println(err)
	}

	for i, e := range list2 {
		fmt.Println(i, ":", e)
	}

}

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
