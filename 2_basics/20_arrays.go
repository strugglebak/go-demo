package main

import (
	"fmt"
)

func main() {
	var a[2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"aaaa",
		"bbbb",
		"cccc",
		"dddd",
	}
	fmt.Println(names)

	// slices are like references to arrays
	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, false},
		{5, true},
	}
	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}
	// [left:right]
	// left: default zero
	// right: the length of the slice
	s3 = s3[1:4]
	fmt.Println(s3)

	s3 = s3[:2]
	fmt.Println(s3)

	s3 = s3[1:]
	fmt.Println(s3)


	// The capacity of a slice is the number of elements in the
	//【underlying array】,
	// counting from the【first element】in the slice.

	// first element is 2
	// so cap is 6
	// len = 6 cap = 6 [2 3 5 7 11 13]
	s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4)

	// first element is 2
	// so cap is 6
	// len = 0 cap = 6 []
	s4 = s4[:0]
	printSlice(s4)

	// first element is 2
	// so cap is 6
	// len = 4 cap = 6 [2 3 5 7]
	s4 = s4[:4]
	printSlice(s4)

	// first element is 5
	// so cap is 4
	// len = 2 cap = 4 [5 7]
	s4 = s4[2:]
	printSlice(s4)

	var s5 []int
	fmt.Println(s5, len(s), cap(s))
	if s5 == nil {
		fmt.Println("nil!")
	}

	// make is using to specify a capacity
	x1 := make([] int, 5)
	printSlice1("x1", x1)

	x2 := make([] int, 0, 5)
	printSlice1("x2", x2)

	x3 := x2[:2]
	printSlice1("x3", x3)

	x4 := x3[2:5]
	printSlice1("x4", x4)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n",  len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s: len = %d cap = %d %v\n",  s, len(x), cap(x), x)
}
