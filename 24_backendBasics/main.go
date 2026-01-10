package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Response struct {
	Error   bool   `json:"error"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
}

// logging middleware
func requestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		fmt.Printf("%s %s %s %s\n",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start))
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := Response{
		Error:   false,
		Data:    nil,
		Message: "api server is live",
	}
	json.NewEncoder(w).Encode(resp)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Println("received: ", u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := Response{
		Error:   false,
		Message: "Success",
		Data:    u,
	}
	json.NewEncoder(w).Encode(resp)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dog := struct {
		Name   string `json:"name"`
		IsGood bool   `json:"good"`
	}{
		"Rex",
		true,
	}

	resp := Response{
		Error:   false,
		Message: "Succes",
		Data:    dog,
	}
	json.NewEncoder(w).Encode(resp)
}

var Port string = ":3000"

func main() {
	http.HandleFunc("/", requestLogger(rootHandler))
	http.HandleFunc("/user", requestLogger(userHandler))
	http.HandleFunc("/random", requestLogger(randomHandler))

	fmt.Println("Server starting at " + Port)

	err := http.ListenAndServe(Port, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}

}
