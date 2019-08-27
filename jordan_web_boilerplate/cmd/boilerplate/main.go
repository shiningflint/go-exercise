package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "response hello world")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8088", nil))
	// srv := &http.Server{
	// 	Addr:         ":8088",
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// 	IdleTimeout:  120 * time.Second,
	// }

	// go func() {
	// 	err := srv.ListenAndServe()
	// 	if err != nil {
	// 		log.Fatal("ListenAndServe:", err)
	// 	}
	// }()
}
