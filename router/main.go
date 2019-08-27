// https://stackoverflow.com/questions/40070168/how-to-build-own-simple-router-in-golang
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// r = newRouter()
	// r.Add("/", hello)
	// fmt.Println("bananas potatoes", r)
	// Run(":8088")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello potato!")
	})
	log.Fatal(http.ListenAndServe(":8088", nil))
}
