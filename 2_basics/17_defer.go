package main

import (
	"fmt"
)

func main() {
	// "world" will be output until main returns
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")
	for i:= 0; i < 10; i++ {
		// defer into a stack
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
