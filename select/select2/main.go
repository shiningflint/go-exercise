package main

import "fmt"

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	fmt.Println("haha")
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}