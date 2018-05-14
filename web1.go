package main

import (
  "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  const content =
`
<h1>Hello world!</h1>
<p>This is a hello world content</p>
<p>This is more hello world content</p>
`
  w.Write([]byte(content))
}

func main() {
  http.HandleFunc("/", indexHandler)
  if err := http.ListenAndServe(":8000", nil); err != nil {
    panic(err)
  }
}
