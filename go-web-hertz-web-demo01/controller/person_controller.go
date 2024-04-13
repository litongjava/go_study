package controller

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/common/utils"
)

type person struct {
  ID   int64  `json:"id,string"` // 注意这里使用了 string 选项
  Name string `json:"name"`
}

func PersonBind(ctx context.Context, c *app.RequestContext) {
  var p person
  if err := c.BindAndValidate(&p); err != nil {
    panic(err)
  }
  c.JSON(200, utils.H{
    "data": p,
    "ok":   true,
  })
}
