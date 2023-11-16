package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world \n")
  })

  fmt.Println("Server is running")
  http.ListenAndServe(":3000", nil)
}
