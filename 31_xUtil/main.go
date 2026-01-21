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

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: x <your query here>")
		return
	}

	userQuery := strings.Join(args, " ")

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

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(userQuery),
		config,
	)
	if err != nil {
		log.Fatal("API error:", err)
	}

	generatedCmd := strings.TrimSpace(result.Text())

	fmt.Printf("x will run: %s\n", generatedCmd)
	fmt.Print("proceed? (y/n): ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice != "y" {
		fmt.Println("Aborted.")
		return
	}

	cmdParts := strings.Fields(generatedCmd)
	if len(cmdParts) == 0 {
		log.Fatal("Generated command was empty")
	}

	out, err := exec.Command(cmdParts[0], cmdParts[1:]...).CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Output: %s", string(out))
		return
	}

	fmt.Println(string(out))
}
