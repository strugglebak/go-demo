package main

import (
	"fmt"
)

// type Vertex struct {
// 	X int
// 	Y int
// }

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	_p = &Vertex{1, 2}
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, _p, v2, v3)
}
