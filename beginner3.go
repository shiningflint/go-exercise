// Modify the beginner2.go. only greet Alice and Bob
package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func inputName() string {
  fmt.Print("What is your name? ")
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  return text
}

func main() {
  name := inputName()
  name = strings.TrimSuffix(name, "\n")
  if name == "Alice" || name == "Bob" {
    fmt.Printf("Hello %s!\n", name)
  } else {
    fmt.Println("No greeting for you!")
  }
}
