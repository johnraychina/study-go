package main

import "fmt"

func main() {

	// panic: assignment to entry in nil map
	//myMap := new(map[string]string)
	//(*myMap)["x"] = "y"
	myMap := make(map[string]string)
	myMap["x"] = "y"
	myMap["u"] = "v"
	fmt.Println(myMap)

	// compare empty struct
	var emp Employee
	if emp == (Employee{}) {
		fmt.Println("empty struct")
	} else {
		fmt.Println(emp)
	}
}

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
