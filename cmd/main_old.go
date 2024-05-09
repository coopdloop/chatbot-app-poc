package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main_old() {
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, "Who was the first man to walk on the moon? Keep the response concise and ask if the answer was what they wanted",
		llms.WithTemperature(0.8),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}
