package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type user struct {
	Name string
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var usr user
	err = json.NewDecoder(body).Decode(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	http.HandleFunc("/hello", helloWorld)
	http.ListenAndServe(port, nil)
}
