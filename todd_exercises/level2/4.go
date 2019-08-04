package main

import (
	"fmt"
)

func main() {
	x := 42000

	fmt.Printf("The var x's value is %v\n", x)
	fmt.Printf("Hex value is %x\n", x)
	fmt.Printf("Binary value is %b\n", x)
	fmt.Printf("Decimal value is %f\n", float64(x))

	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}
