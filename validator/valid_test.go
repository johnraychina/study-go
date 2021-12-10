package validator

import "testing"

//import "github.com/go-playground/validator/v10"

type MyStruct struct {
	name string ``
}

func TestName(t *testing.T) {
	//err := validate.struct(MyStruct)
	//validationErrors := err.(validator.ValidationErrors)
}
