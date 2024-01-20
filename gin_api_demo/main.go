package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

type Micheal map[string]interface{}

func pong(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "pong",
  })
}
func hello(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "hello func",
  })
}
func test(c *gin.Context) {
  c.JSON(http.StatusOK, map[string]interface{}{
    "message": "自己写的map没有使用gin.H",
    "code":    200,
  })
}
func micheal(c *gin.Context) {
  c.JSON(http.StatusOK, Micheal{
    "message": "使用我们自己定义的Micheal类型",
    "code":    200,
    "data":    "",
  })
}
func main() {
  // Create an instance of Engine, by using New() or Default()
  // 创建一个gin默认的Engine实例
  r := gin.Default()
  // 用户使用get方式去请求地址为/ping的时候执行pong函数
  r.GET("/ping", pong)
  r.GET("/hello", hello)
  r.GET("/test", test)
  r.GET("/micheal", micheal)
  // iris.map
  // gin.H
  // 启动 默认端口8080
  // 如果需要修改端口 r.Run(":1234")
  r.Run(":1234")
}
