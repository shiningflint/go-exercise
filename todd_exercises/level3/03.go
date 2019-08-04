package main

import (
	"fmt"
)

func main() {
	for i := 65; i < 91; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%U %c\n", i, i)
		}
	}
}
