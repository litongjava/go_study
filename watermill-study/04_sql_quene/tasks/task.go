package tasks

import (
  "context"
  "encoding/json"
  "github.com/ThreeDotsLabs/watermill/message"
)

type TaskTopic string

const (
  MessageSummarizerTopic      TaskTopic = "message_summarizer"
  MessageEmbedderTopic        TaskTopic = "message_embedder"
  MessageNerTopic             TaskTopic = "message_ner"
  MessageIntentTopic          TaskTopic = "message_intent"
  MessageTokenCountTopic      TaskTopic = "message_token_count"
  DocumentEmbedderTopic       TaskTopic = "document_embedder"
  MessageSummaryEmbedderTopic TaskTopic = "message_summary_embedder"
  MessageSummaryNERTopic      TaskTopic = "message_summary_ner"
)

type Task interface {
  Execute(ctx context.Context, event *message.Message) error
  HandleError(err error)
}

// SimpleTask 实现 Task 接口用于测试
type SimpleTask struct {
  Processed bool
}

func (t *SimpleTask) Execute(ctx context.Context, msg *message.Message) error {
  marshal, err := json.Marshal(msg)
  if err != nil {
    llog.Error(err.Error())
  } else {
    llog.Info("msg,", string(marshal))
  }

  t.Processed = true
  return nil
}

func (t *SimpleTask) HandleError(err error) {
  // 处理消息处理过程中的任何错误
}

func Initialize(ctx context.Context, appState *AppState, router *TaskRouter) {
  llog.Info("Initializing tasks")
}
