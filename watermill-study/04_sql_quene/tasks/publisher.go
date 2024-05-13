package tasks

import (
  "database/sql"
  "encoding/json"
  "fmt"
  "github.com/ThreeDotsLabs/watermill"
  "github.com/ThreeDotsLabs/watermill/message"
  wla "github.com/ma-hartma/watermill-logrus-adapter"
)

type TaskPublisher struct {
  publisher message.Publisher
}

func NewTaskPublisher(db *sql.DB) *TaskPublisher {
  var wlog = wla.NewLogrusLogger(llog)
  publisher, err := NewSQLQueuePublisher(db, wlog)
  if err != nil {
    llog.Fatalf("Failed to create task publisher: %v", err)
  }
  return &TaskPublisher{
    publisher: publisher,
  }
}

// Publish publishes a message to the given topic. Payload must be a struct that can be marshalled to JSON.
func (t *TaskPublisher) Publish(
  taskType TaskTopic,
  metadata map[string]string,
  payload any,
) error {
  llog.Debugf("Publishing task: %s", taskType)
  p, err := json.Marshal(payload)
  if err != nil {
    return fmt.Errorf("failed to marshal message: %w", err)
  }
  llog.Debugf("Publishing message: %s", p)
  m := message.NewMessage(watermill.NewUUID(), p)
  m.Metadata = metadata

  err = t.publisher.Publish(string(taskType), m)
  if err != nil {
    return fmt.Errorf("failed to publish task message: %w", err)
  }

  llog.Debugf("Published task: %s", taskType)

  return nil
}

func (t *TaskPublisher) Close() error {
  err := t.publisher.Close()
  if err != nil {
    return fmt.Errorf("failed to close task publisher: %w", err)
  }

  llog.Debug("Closed task publisher")

  return nil
}
