package main

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // GET 获取所有的文章信息
  r.GET("/posts", func(c *gin.Context) {
    c.String(http.StatusOK, "GET")
  })
  // POST 创建一篇新文章
  r.POST("/posts", func(c *gin.Context) {
    c.String(http.StatusOK, "POST")
  })
  // PUT 修改一篇文章
  r.PUT("/posts/:id", func(c *gin.Context) {
    c.String(http.StatusOK, fmt.Sprintf("PUT id: %s", c.Param("id")))
  })
  // DELETE 删除一篇文件
  r.DELETE("/posts/", func(c *gin.Context) {
    c.String(http.StatusOK, "DELETE")
  })

  // 匹配所有请求方法
  /*
  	MethodGet     = "GET"
  	MethodHead    = "HEAD"
  	MethodPost    = "POST"
  	MethodPut     = "PUT"
  	MethodPatch   = "PATCH" // RFC 5789
  	MethodDelete  = "DELETE"
  	MethodConnect = "CONNECT"
  	MethodOptions = "OPTIONS"
  	MethodTrace   = "TRACE"
  */
  r.Any("/users", func(c *gin.Context) {
    c.String(200, "any")
  })

  r.Run()
}
