package main

import (
	"fmt"
	"net/http"
	"time"
)

func requestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		fmt.Printf("\n%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start))
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server")
}

var Port string = ":3000"

func main() {
	fmt.Println("hellow world")
	http.HandleFunc("/", requestLogger(root))
	fmt.Println("Server starting at " + Port)

	err := http.ListenAndServe(Port, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}

}
