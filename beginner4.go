// get number input n, and then print the sum of 1 to n
package main

import (
  "fmt"
  "os"
  "bufio"
)

func readInt(prompt string) int {
  fmt.Println(prompt)

  reader := bufio.NewReader(os.Stdin)

  var input int
  fmt.Fscan(reader, &input)
  return input

}

func sumOneToNumber(number int) int {
  const numberOne int = 1
  var result int = numberOne

  for i := numberOne; i < number; i++ {
    result = result + (i + 1)
  }

  return result
}

func main() {
  fmt.Println("this is beginner 4")
  number := readInt("Input a number")
  result := sumOneToNumber(number)
  fmt.Println(result)
}
