package reflection

import (
	"fmt"
	"reflect"
)

/*
* write walk(x interface{}, cb func(string))
* that calls cb for every string in x where x is a struct
 */

type Ciccio struct {
	string_one string
	number_one int
	string_two string
	number_two int
}

var pippo = Ciccio{
	"string_one",
	2,
	"string_two",
	4,
}

func walk(x any, cb func(s string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Println(field.String())
		cb(field.String())
	}
}
