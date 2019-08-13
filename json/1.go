package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	unMarshalExample()
	marshalExample()
}

func unMarshalExample() {
	a := []byte(`{"bananas": "person"}`)
	fmt.Println(string(a))
	var x map[string]interface{}
	if err := json.Unmarshal(a, &x); err != nil {
		fmt.Println(err)
	}
	fmt.Println(x["bananas"])
}

func marshalExample() {
	type person struct {
		Fname string `json:"firstName"`
		Lname string `json:"lastName"`
	}

	a := person{Fname: "Adum", Lname: "Keristoper"}
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
