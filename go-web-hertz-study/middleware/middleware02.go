package main

import (
  "context"
  "errors"
  "github.com/bytedance/gopkg/util/logger"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/app/server"
  "github.com/cloudwego/hertz/pkg/common/hlog"
  "github.com/cloudwego/hertz/pkg/common/utils"
)

func main() {
  h := server.Default()
  h.Use(GlobalErrorHandler)
  //获取的是Query参数，即？后面的参数
  h.GET("/user", func(ctx context.Context, c *app.RequestContext) {
    name := c.Query("name")
    c.String(200, name)
  })
  h.Spin()
}
func GlobalErrorHandler(ctx context.Context, c *app.RequestContext) {
  c.Next(ctx)

  if len(c.Errors) == 0 {
    // 没有收集到异常直接返回
    hlog.Info("no error")
    return
  } else {
    hlog.Info("error occurred")
  }
  hertzErr := c.Errors[0]
  // 获取errors包装的err
  err := hertzErr.Unwrap()
  // 打印异常堆栈
  logger.CtxErrorf(ctx, "%+v", err)
  // 获取原始err
  err = errors.Unwrap(err)
  // todo 进行错误类型判断
  c.JSON(400, utils.H{
    "code":    400,
    "message": err.Error(),
  })
}
