package main

import (
	"fmt"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/ request")
	fmt.Fprintf(w, "Ola, Docker")
}

func main() {
	http.HandleFunc("/", getRoot)
	fmt.Printf("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}
