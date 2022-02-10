package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <- tick:
			fmt.Println("tick")
		case <- boom:
			fmt.Println("BOOM!")
			return
		// The default case in a select is run if no other case is ready.
		default:
			fmt.Print("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// lets a goroutine wait on multiple communication operations
		// A select blocks until one of its cases can run, then it executes that case.
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}
