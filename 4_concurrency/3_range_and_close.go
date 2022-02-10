package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	for i := range c {
		fmt.Print(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
