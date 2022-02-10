package main

import "fmt"

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	// chan means channel
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// receive from c
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// Channels can be buffered.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// send sum to c
	c <- sum
}
