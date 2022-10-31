package test

import (
	"fmt"
	"reflect"
)

type String optionalString
type optionalString []string

func (o String) Digest() bool {
	//var emptyVal int
	//emptyVal = 233
	//reflect.
	//	fmt.Println("check o", reflect.ValueOf(o).Interface())

	switch reflect.ValueOf(o).Interface().(type) {
	case string, []byte:
		fmt.Println("ok.this is string ")
	case bool:
		fmt.Println("ok this is bool")
	case int:
		fmt.Println("ok this. is int ")
	case []string:
		fmt.Println("this is []string")

	}

	fmt.Println("check....value")

	return true

}
