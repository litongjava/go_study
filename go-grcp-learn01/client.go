package main

import (
  "context"
  "fmt"
  "go-grcp-learn01/message"
  "google.golang.org/grpc"
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

  //调用远程服务
  orderRequest := &message.OrderRequest{OrderId: "201907300001", TimeStamp: time.Now().Unix()}
  orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
  //输出结果
  if orderInfo != nil {
    fmt.Println(orderInfo.GetOrderId())
    fmt.Println(orderInfo.GetOrderName())
    fmt.Println(orderInfo.GetOrderStatus())
  }
}
