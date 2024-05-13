package tasks

import (
  "context"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestPublishAndSubscribeMessage(t *testing.T) {
  // 初始化数据库连接
  //var dsn = "postgres://postgres:AVNS_EOZHj-v6YwugeWM8P2m@127.0.0.1:15432/defaultdb?sslmode=disable"
  var dsn = "postgres://postgres:123456@192.168.3.9:5432/postgres?sslmode=disable"
  db, err := NewPostgresConnForQueue(dsn)
  assert.NoError(t, err)
  defer db.Close()

  // 初始化应用状态
  appState := &AppState{}
  // run the router
  RunTaskRouter(context.Background(), appState, db)

  // check that the router is configured
  assert.NotNil(t, appState.TaskRouter, "task router is nil")
  assert.NotNil(t, appState.TaskPublisher, "task publisher is nil")

  //消息通道
  //taskType := MessageSummarizerTopic
  taskType := MessageEmbedderTopic

  // 定义并添加任务处理器
  testTask := &SimpleTask{}
  appState.TaskRouter.AddTask(context.Background(), "test_task_handler", taskType, testTask)

  // 定义消息类型和内容
  metadata := map[string]string{"key1": "value1"}
  payload := struct {
    Data string
  }{
    Data: "example payload",
  }

  // 发布消息
  err = appState.TaskPublisher.Publish(taskType, metadata, payload)
  assert.NoError(t, err, "failed to publish message")

  llog.Info("running tasks router")
  err = appState.TaskRouter.Run(context.Background())
  if err != nil {
    llog.Fatalf("failed to run tasks router %v", err)
  }
  defer appState.TaskRouter.Close()
  defer appState.TaskPublisher.Close()
}
