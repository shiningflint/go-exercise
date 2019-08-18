package main

import "fmt"

func main() {
	go foo()
	bar()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar:", i)
	}
}
