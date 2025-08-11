package main

import (
	"context"

	"github.com/openai/openai-go/v2"
)

func main() {
	client := openai.NewClient(
	// option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)
	msgs := []openai.ChatCompletionMessageParamUnion{
		openai.UserMessage("Tell me the cheapest SLM as of 2025"),
		openai.SystemMessage("You are an expert machine learning engineer at Google."),
	}

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: msgs,
		Model:    openai.ChatModelGPT4_1,
	})
	if err != nil {
		panic(err.Error())
	}

	println(chatCompletion.RawJSON())
	println(chatCompletion.Choices[0].Message.Content)
}
