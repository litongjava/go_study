package main

import (
  "context"
  "fmt"
  "go-grpc-server-stream-learn/message"
  "google.golang.org/grpc"
  "io"
  "time"
)

func main() {
  //1、Dail连接
  conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
  if err != nil {
    panic(err.Error())
  }
  defer conn.Close()

  //创建client
  orderServiceClient := message.NewOrderServiceClient(conn)

  //发送请求
  request := message.OrderRequest{TimeStamp: time.Now().Unix()}

  orderInfosClient, err := orderServiceClient.GetOrderInfos(context.TODO(), &request)

  //接收数据并输出
  for {
    orderInfo, err := orderInfosClient.Recv()
    if err == io.EOF {
      fmt.Println("读取结束")
      return
    }
    if err != nil {
      panic(err.Error())
    }
    fmt.Println("读取到的信息：", orderInfo)
  }
}
