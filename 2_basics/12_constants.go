package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	const World = "world"
	fmt.Print("hello", World)
	fmt.Println("happy", Pi, "day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
