package main

import (
  "fmt"
  "go-rcp-learn/param"
  "net/rpc"
)

//客户端逻辑实现
func main() {

  client, err := rpc.DialHTTP("tcp", "localhost:8081")
  if err != nil {
    panic(err.Error())
  }

  var req float32 //请求值
  req = 5

  //var resp *float32 //返回值
  //同步的调用方式
  //err = client.Call("MathUtil.CalculateCircleArea", req, &resp)
  //if err != nil {
  //	panic(err.Error())
  //}
  //fmt.Println(*resp)

  var respSync *float32
  //异步的调用方式
  syncCall := client.Go("MathUtil.CalculateCircleArea", req, &respSync, nil)
  //fmt.Println(*respSync)
  replayDone := <-syncCall.Done
  fmt.Println(replayDone)
  fmt.Println(*respSync)

  var result *float32
  addParma := &param.AddParma{Args1: 1.2, Args2: 2.3}
  err = client.Call("MathUtil.Add", addParma, &result)
  if err != nil {
    panic(err.Error())
  }
  fmt.Println("计算结果:", *result)
}
