package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a simple golang http webserver.")
}

var Port string = ":3000"

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting at :3000")
	fmt.Println("routes: /, /about")

	err := http.ListenAndServe(Port, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
	// any code here will not be executed as http.ListenAndServe blocks for listening mode
	fmt.Println("test string")
}
