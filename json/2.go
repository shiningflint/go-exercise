package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	enc := json.NewEncoder(os.Stdout)
	fmt.Printf("%T\n", enc)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
