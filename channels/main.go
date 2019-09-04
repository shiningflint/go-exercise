package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("bananas")
	n := 3
	in := make(chan int)
	out := make(chan int)

	// We now supply 2 channels to the `multiplyByTwo` function
	// One for sending data and one for receiving
	go multiplyByTwo(in, out)

	// We then send it data through the channel and wait for the result
	go setInt(n, in)
	fmt.Println(<-out)
}

func setInt(n int, in chan<- int) {
	time.Sleep(time.Second)
	fmt.Println("going to put the n which is 3")
	in <- n
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	// This line is just to illustrate that there is code that is
	// executed before we have to wait on the `in` channel
	fmt.Println("init go routine...")

	// The goroutine does not proceed until data is received on the `in` channel
	fmt.Println("wating for some in from somewhere")
	num := <-in

	// The rest is unchanged
	result := num * 2
	fmt.Println("I'm finished, sending the result")
	out <- result
}
