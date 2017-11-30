package main

import (
  "fmt"
  "net/http"
  "os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Request received")
  fmt.Fprintf(w, "Hello world!\n")
}

func main() {

  port := os.Getenv("HTTP_PORT")

  if port == "" {
    port = ":9000"
  }

  fmt.Println("Started server")
  fmt.Println("Serving on port", port)
  http.HandleFunc("/hello", HelloWorld)
  http.ListenAndServe(port, nil)
}
