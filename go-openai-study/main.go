package main

import (
  "context"
  "github.com/sashabaranov/go-openai"
  "os"
)

import (
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }

  client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
  request := openai.ChatCompletionRequest{
    //Model: "gpt-4-gizmo-g-k3IqoCe1l",
    //Model: "gpt-4-gizmo-g-k3IqoCe1l-code-guru",
    //Model: "gpt-4-gizmo-g-uBhKUJJTl",
    Model: openai.GPT3Dot5Turbo,
    Messages: []openai.ChatCompletionMessage{
      {
        Role:    openai.ChatMessageRoleSystem,
        Content: "you are a helpful chatbot",
      },
      {
        Role:    openai.ChatMessageRoleUser,
        Content: "Can you review this code snippet?",
      },
    },
  }
  response, err := client.CreateChatCompletion(context.Background(), request)
  if err != nil {
    print(err.Error())
  } else {
    // 解析响应数据
    result := response.Choices[0].Message.Content //回答的内容
    print(result)
  }
}
