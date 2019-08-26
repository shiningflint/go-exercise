// https://stackoverflow.com/questions/40070168/how-to-build-own-simple-router-in-golang
package main

import (
	"fmt"
	"net/http"
)

// Handle type is the http handler function
type Handle func(http.ResponseWriter, *http.Request)

// Router type has mux which is a map with string keys
// and values are type Handle
type Router struct {
	mux map[string]Handle
}

// Add method will add a route to the mux path
func (r *Router) Add(path string, handle Handle) {
	r.mux[path] = handle
}

func newRouter() *Router {
	return &Router{
		mux: make(map[string]Handle),
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world bananas!")
}

func main() {
	r := newRouter()
	fmt.Println("bananas potatoes", r)
}
