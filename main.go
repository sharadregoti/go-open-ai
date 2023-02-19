package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading input: %v", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		return
	}

	c := gogpt.NewClient(apiKey)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		Prompt:           string(input),
		Temperature:      0,
		MaxTokens:        256,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Stream:           true,
	}
	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		return
	}

	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}

		if len(response.Choices) == 0 {
			continue
		}

		fmt.Printf("%v", response.Choices[0].Text)
	}
}
