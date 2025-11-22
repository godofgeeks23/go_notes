package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const myURL = "https://backend.reconaid.in"
const customURL = "https://localhost:3000/posts?filter=today&sortby=recent"

func main() {
	res, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	// fmt.Println(res)
	fmt.Printf("type of http response - %T\n", res)

	dataBytes, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("GET response - ", string(dataBytes))

	// url parsing
	result, err := url.Parse(customURL)
	if err != nil {
		fmt.Println("error in url parsing -", err)
		panic(err)
	}
	fmt.Printf("scheme - %v\nhost - %v\npath - %v\nport - %v\nraw query - %v\n", result.Scheme, result.Host, result.Path, result.Port(), result.RawQuery)
}
