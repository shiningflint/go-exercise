package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	// n := "53"
	// fmt.Println("bananas")
	// fmt.Println(parseInt(n))
	// fmt.Printf("%T\n", parseInt(n))

	a := "Adum"
	fmt.Println(a)
	a = "banana"
	fmt.Println(a)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func parseInt(value string) int {
	result, err := strconv.Atoi(value)
	if err == nil {
		return result
	} else {
		return 0
	}
}
