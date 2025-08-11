package main

import (
	"context"

	"github.com/openai/openai-go/v2"
)

func main() {
	client := openai.NewClient(
		// option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("What is the importance of using small model languages vs large model languages as of 2025"),
			openai.SystemMessage("You are an expert machine learning engineer at Google."),			
		},
		Model: openai.ChatModelGPT5,
	})
	if err != nil {
		panic(err.Error())
	}
	println(chatCompletion.Choices[0].Message.Content)
}
