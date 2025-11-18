package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	dir := os.Args[1]
	port := ":" + os.Args[2]

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	fmt.Printf("Launching a file server at %v/  Access on port %v", dir, port)
	http.ListenAndServe(port, nil)
}
