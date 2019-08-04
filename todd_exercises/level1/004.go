package main

import (
	"fmt"
)

type hooman int

var x hooman

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
