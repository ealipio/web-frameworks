package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", handler)
	// run server on port 8080
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
