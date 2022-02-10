package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Stringer defined by the fmt package.
// type Stringer interface {
//     String() string
// }
// The fmt package (and many others) look for this interface to print values
func (p Person) String() string {
	return fmt.Sprintf("%v: (%v years old now)", p.Name, p.Age)
}

func main() {
	a := Person{"strugglebak", 42}
	b := Person{"strugglebak2", 422}
	fmt.Println(a, b)
}
