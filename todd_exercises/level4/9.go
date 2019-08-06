package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	x[`mumu`] = []string{`lala`, `bananas`, `potatoes`}

	for k, v := range x {
		fmt.Printf("%s: ", k)
		for i, xs := range v {
			fmt.Printf("%s", xs)
			if i != len(v)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")
	}
}
