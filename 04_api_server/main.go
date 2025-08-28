package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Simple response structure for JSON output
type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "This is a simple Go API webserver."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Example: Just return a fixed user JSON
	user := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "John Doe",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/user", userHandler)

	fmt.Println("Starting server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
