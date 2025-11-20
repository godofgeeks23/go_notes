package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const myURL = "https://backend.reconaid.in/"

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
	fmt.Println("response - ", string(dataBytes))

	result, _ := url.Parse(myURL)
	fmt.Println(result.Scheme, result.Host, result.Path, result.Port(), result.RawQuery)

}
