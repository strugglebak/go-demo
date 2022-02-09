package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2077

	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Print(i)

	fmt.Printf("\n")

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
