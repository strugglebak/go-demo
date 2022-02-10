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

	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
