package main

import "fmt"

func main() {
	fmt.Println("today I'm going to try some channels")
	messages := make(chan string)

	go func() {
		messages <- "banana"
	}()

	go func() {
		messages <- "potato"
	}()

	a, b := <-messages, <-messages
	fmt.Println(a, b, a+b)
}
