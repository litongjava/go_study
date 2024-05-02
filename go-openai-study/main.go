package main

import (
	"context"
	"os"
)

import (
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	request := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "你好",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: "你好，有什么可以帮助你的嘛？",
			},
		},
	}
	response, err := client.CreateChatCompletion(context.Background(), request)
	if err != nil {
		// 错误处理
	} else {
		// 解析响应数据
		result := response.Choices[0].Message.Content //回答的内容
		print(result)
	}
}
