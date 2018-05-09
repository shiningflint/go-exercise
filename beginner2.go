// Write a program that asks the user for their name and greet them
package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func inputName(greeting string) string {
  fmt.Print(greeting)

  // Creates a reader and reads the string ended by end of line \n so it accepts spaces
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')

  return text
}

func main() {
  name := inputName("Hi! What is your name? ")

  // remove the \n inputted after the enter
  name = strings.TrimSuffix(name, "\n")
  fmt.Printf("Hello %s!\n", name)
}
