package main

import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// An interface type is defined as a set of method signatures.
type Abser interface {
	Abs() float64
}

type I interface {
	M()
}
type T struct {
	S string
}
// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// Interfaces are implemented implicitly
// no "implements" keyword.
// func (t T) M() {
// 	fmt.Println(t.S)
// }
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	// var i I = T{"hello"}
	// i.M()

	var j I
	j = &T{"hello J"}
	describe(j)
	j.M()

	j = F(math.Pi)
	describe(j)
	j.M()

	var k I
	var t *T
	k = t
	// k is nil
	describe(k)
	k.M()

	k = &T{"hello K"}
	// k is not nil
	describe(k)
	k.M()

	// runtime error
	// there is no type inside the interface tuple
	// to indicate which concrete method to call.
	// var z I
	// describe(z)
	// z.M()

	var z interface{}
	describe2(z)

	z = 32
	describe2(z)

	z = "hello"
	describe2(z)
}

func describe(i I) {
	// interface values
	// (value, type)
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe2(i interface{}) {
	// interface values
	// (value, type)
	fmt.Printf("(%v, %T)\n", i, i)
}
