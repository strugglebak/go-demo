package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// define methods on types
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Scale(v *Vertex, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


// declare a method on non-struct types
type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	ScaleFunc(&v, 10)
	fmt.Println(v.Abs())

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v, p)
	fmt.Println(AbsFunc(*p))

	Scale(&v, 10)
	fmt.Println(Abs(v))

	a := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", a, a.Abs2())
	a.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", a, a.Abs2())
}
