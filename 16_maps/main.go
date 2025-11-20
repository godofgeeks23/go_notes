package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps")

	var langs = make(map[string]string)
	langs["js"] = "javascript"
	langs["go"] = "golang"
	langs["py"] = "python"
	fmt.Println(langs)
	fmt.Println(langs["js"])
	delete(langs, "js")
	fmt.Println(langs)

	for key, value := range langs {
		fmt.Println(key, value)
	}
}
