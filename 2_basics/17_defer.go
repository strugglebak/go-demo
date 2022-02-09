package main

import (
	"fmt"
)

func main() {
	// "world" will be output until main returns
	defer fmt.Println("world")
	fmt.Println("hello")
}
