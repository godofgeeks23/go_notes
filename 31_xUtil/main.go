package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	fmt.Println("x (x-utility v1.0.0)")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error reading .env file")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("API key not found")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	systemPrompt := "you are a command generator for a linux based os. user will submit a action that he wants to perform. generate the command needed to run to do the action. return only the command - nothing else (no explanation, no comments or anything - just the required command)"

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(systemPrompt, genai.RoleUser),
	}

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text("remove unused docker volumes"),
		config,
	)

	fmt.Println("x will run the following command on approval:", result.Text())
	fmt.Println("proceed?(y/n)")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	if choice != "y" {
		log.Fatal("aborted")
	}
	out, err := exec.Command(result.Text()).Output()
	if err != nil {
		log.Fatal("error running command")
		fmt.Printf("%s\n", err)
	}
	output := string(out[:])
	fmt.Println(output)

}
