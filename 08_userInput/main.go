package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter username - ")

	// comma, ok / comma, error syntax
	name, _ := reader.ReadString('\n')
	// can be name, err too. it is like try catch. err will catch the errors

	welcome := "welcome to golang " + name
	fmt.Println(welcome)

	fmt.Println("enter a number between 1 and 10")
	ratingString, _ := reader.ReadString('\n')
	ratingNum, err := strconv.ParseFloat(strings.TrimSpace(ratingString), 64)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("wow! thanks for rating+1 = ", ratingNum+1)
	}
}
