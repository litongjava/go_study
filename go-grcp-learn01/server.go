package main

import (
  context "context"
  "errors"
  "fmt"
  "go-grcp-learn01/message"
  "google.golang.org/grpc"
  "log"
  "net"
  "time"
)

type OrderServiceImpl struct {
}

//具体的方法实现
func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {
  orderMap := map[string]message.OrderInfo{
    "201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
    "201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
    "201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
  }

  var response *message.OrderInfo
  current := time.Now().Unix()
  if request.TimeStamp > current {
    *response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
  } else {
    result := orderMap[request.OrderId]
    if result.OrderId != "" {
      fmt.Println(result)
      return &result, nil
    } else {
      return nil, errors.New("server error")
    }
  }
  return response, nil
}

func main() {
  start := time.Now().UnixMilli()

  server := grpc.NewServer()

  message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

  lis, err := net.Listen("tcp", ":8090")
  if err != nil {
    panic(err.Error())
  } else {
    log.Println("listen :8090")
  }
  end := time.Now().UnixMilli()
  log.Printf("%d(ms)\n", end-start)
  server.Serve(lis)

}
