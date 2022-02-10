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
func (t T) M() {
	fmt.Println(t.S)
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

	var i I = T{"hello"}
	i.M()
}
