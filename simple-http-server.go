package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	Name string `json:"Name"`
}

// for health checks
func CheckHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// handle simple json request
func HelloWorld(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "Request body is missing", http.StatusBadRequest)
		return
	}

	var usr User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reply := fmt.Sprintf("Hello %s", usr.Name)
	fmt.Fprintf(w, reply)
}

func main() {

	port := os.Getenv("HTTP_PORT")

	if port == "" {
		port = ":9000"
	}

	fmt.Println("Started server")
	fmt.Println("Serving on port", port)
	http.HandleFunc("/hello", HelloWorld)
	http.HandleFunc("/health", CheckHealth)
	http.ListenAndServe(port, nil)
}
