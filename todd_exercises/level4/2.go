package main

import "fmt"

func main() {
	x := []int{11, 66, 33, 66, 77, 92, 2, 5, 6, 7, 8}
	for _, v := range x {
		fmt.Printf("%d - %T\n", v, v)
	}
	fmt.Printf("%T\n", x)
}
