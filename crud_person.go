package main

import (
	"bufio"
	"fmt"
	"os"
)

type person struct {
	fname string
	lname string
}

func main() {
	var people []person

	sysInput := ""
	for sysInput != "exit" {
		menu()
		sysInput = readString("Your Input: ")
		if sysInput == "1" {
			indexPerson(people)
		} else if sysInput == "2" {
			people = append(people, newPerson(people))
		} else {
			sysInput = "exit"
		}
	}
}

func readString(prompt string) string {
	fmt.Printf("%s", prompt)

	reader := bufio.NewReader(os.Stdin)

	var input string
	fmt.Fscan(reader, &input)
	return input
}

func menu() {
	fmt.Printf(`
Menu:
1. List of registered people
2. Register new person
`)
}

func indexPerson(people []person) {
	fmt.Println(people)
}

func newPerson(people []person) person {
	fmt.Println("Let's create a new person!")
	var fname string = readString("Persons First Name: ")
	var lname string = readString("Persons Last Name: ")
	fmt.Println("success!")
	return person{fname: fname, lname: lname}
}
