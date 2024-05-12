package main

import (
  "context"
  "fmt"
  "github.com/ThreeDotsLabs/watermill"
  "github.com/ThreeDotsLabs/watermill/message"
  "github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
  "time"
)

func main() {
  pubSub := gochannel.NewGoChannel(
    gochannel.Config{},
    watermill.NewStdLogger(false, false),
  )

  messages, err := pubSub.Subscribe(context.Background(), "example.topic")
  if err != nil {
    panic(err)
  }

  go process(messages)
  publishMessages(pubSub)
}

func process(messages <-chan *message.Message) {
  for msg := range messages {
    fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))
    //我们需要确认我们已经接收并处理了消息，否则它将被重发多次。
    msg.Ack()
  }
}

func publishMessages(publisher message.Publisher) {
  for {
    msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))

    if err := publisher.Publish("example.topic", msg); err != nil {
      panic(err)
    }

    time.Sleep(time.Second)
  }
}
