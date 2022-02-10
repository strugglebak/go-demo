package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// A type assertion provides access to an interface value's underlying concrete value.
	// is interface value i a string type?
	// yes
	s := i.(string)
	fmt.Println(s)

	// is interface value i a string type?
	// yes
	// ok is true
	s, ok := i.(string)
	fmt.Println(s, ok)

	// is interface value i a float64 type?
	// no
	// ok is false
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic: interface conversion: interface {} is string, not float64
	// f = i.(float64)
	// fmt.Println(f)
}
