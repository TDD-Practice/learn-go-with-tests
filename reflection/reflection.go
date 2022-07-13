package reflection

import "fmt"

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

func walk(x any, cb func(string)) {
	fmt.Println("walking the talk")
}
