package main

import "fmt"

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{a, b}

	for _, v := range x {
		for _, k := range v {
			fmt.Println(k)
		}
	}
}
