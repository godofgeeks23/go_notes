package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func main() {
	example := Response{Error: false, Message: "No error!"}

	// marshalling - go obj to json
	jsonData, err := json.Marshal(example)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(example)
	fmt.Println(jsonData)
	fmt.Println(string(jsonData))

	jsonContent := `{
	"error": true,
	"message": "An error occured"
	}`
	var newResponse Response
	err = json.Unmarshal([]byte(jsonContent), &newResponse)
	if err != nil {
		fmt.Println("an error occured while unmarshalling")
	}
	fmt.Println(newResponse)
}
