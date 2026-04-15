package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	l := len(os.Args)
	if l < 3 {
		fmt.Printf("\n\tNeed Args: model prompt\n\n")
		return
	}
	model := os.Args[1]
	prompt := strings.Join(os.Args[2:], " ")

	// Initialize the Ollama client
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Generate a response
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nResponse from %s:\n\n%s", model, completion)
}
