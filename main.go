package main

import (
	"fmt"
	
	"net/http"
)

func main() {
  fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

  http.HandleFunc("/upload", uploadHandler)

  fmt.Println("Server is running")
  http.ListenAndServe(":3000", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint upload hit")
  r.ParseMultipartForm(10 << 20)

  file, handle, err := r.FormFile("myFile")

  if err != nil {
    fmt.Println("Failed to retrieving file:", err)
    return
  }

  defer file.Close()

  fmt.Println("Uploaded file:", handle.Filename)
  fmt.Println("Size:", handle.Size)
  fmt.Println("MIME Header:", handle.Header)


  fmt.Fprint(w, "Upload is complete")

}
