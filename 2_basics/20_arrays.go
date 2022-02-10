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
}
