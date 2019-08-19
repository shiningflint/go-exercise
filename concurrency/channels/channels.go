package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("today I'm going to try some channels")
	messages := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		messages <- "banana"
	}()

	go func() {
		messages <- "potato"
	}()

	a := <-messages
	fmt.Println(a)

	b := <-messages
	fmt.Println(b)

	fmt.Println(a + b)
}
