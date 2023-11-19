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

  http.HandleFunc("/upload", uploadHandler)

  fmt.Println("Server is running")
  http.ListenAndServe(":3000", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Uploading file")
}
