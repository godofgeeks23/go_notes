package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files")

	fileName := "./sample.txt"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}

	length, err := io.WriteString(file, "a sample line for a sample file")
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	fmt.Println("length written:", length)
	defer file.Close()

	// reading a file
	dataByte, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	fmt.Println(dataByte)
	fmt.Println(string(dataByte))
}
