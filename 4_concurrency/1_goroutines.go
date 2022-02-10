package main

import (
	"fmt"
	"time"
)

func main() {
	// happens in the new goroutine
	go say("world")
	// happens in the current goroutine
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}
