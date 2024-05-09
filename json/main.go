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
	//var list = []Employee{
	//	{Name: "Jack", Age: 22},
	//	{Name: "Patrick", Age: 22},
	//}
	//bytes, err := json.Marshal(list)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//list2 := make([]*Employee, 0)
	//err = json.Unmarshal(bytes, &list2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for i, e := range list2 {
	//	fmt.Println(i, ":", e)
	//}
	//
	//remark := ""
	//if strings.Contains(remark, "init by migration") {
	//	panic("init by migration")
	//}

	//extract()

	emp := &Employee{Name: "Kevin", Age: 18}
	fmt.Printf("%+v \n", emp)

	json.Unmarshal([]byte(`{"Age":20}`), &emp)
	fmt.Printf("%+v \n", emp)

	json.Unmarshal([]byte(`{}`), &emp)
	fmt.Printf("%+v \n", emp)

	json.Unmarshal([]byte(`{"FakeField":"fake"}`), &emp)
	fmt.Printf("%+v \n", emp)
}

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func extract() {
	jsonStr := `{"remark":"cold_data", "other":{"key": 1} }`

	var data map[string]string
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for key, value := range data {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
