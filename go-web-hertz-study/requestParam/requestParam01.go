package main

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/app/server"
  "github.com/cloudwego/hertz/pkg/common/utils"
)

func main() {
  h := server.Default()
  h.POST("/person_bind/:age", PersonBind)
  h.Spin()
}

type person struct {
  Age  int    `path:"age" json:"age"`    // 从路径中获取参数
  Name string `query:"name" json:"name"` // 从query中获取参数
  City string `json:"city"`              // 从body中获取参数
}

func PersonBind(ctx context.Context, c *app.RequestContext) {
  var p person
  if err := c.BindAndValidate(&p); err != nil {
    panic(err)
  }
  c.JSON(200, utils.H{
    "person": p,
  })
}
