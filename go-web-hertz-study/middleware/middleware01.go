package main

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/app/server"
  "github.com/hertz-contrib/cors"
  "time"
)

func main() {
  h := server.Default()

  //任何不符合匹配规则的uri请求都会返回空
  h.NoRoute(func(c context.Context, ctx *app.RequestContext) {
    println(ctx.FullPath())
  })
  //跨域中间件
  //app.HandlerFunc
  corsFunction := cors.New(cors.Config{
    AllowAllOrigins:  true,                                     //允许所有origin的请求
    AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"}, //允许的方法
    AllowHeaders:     []string{"Origin"},                       //允许的头部
    ExposeHeaders:    []string{"Content-Length"},               //暴漏的头部信息
    AllowCredentials: true,                                     //允许携带证书
    AllowWildcard:    true,                                     //允许使用通配符匹配
    MaxAge:           12 * time.Hour,                           //请求缓存的最长时间
  })

  h.Use(corsFunction)

  //获取的是Query参数，即？后面的参数
  h.GET("/user", func(ctx context.Context, c *app.RequestContext) {
    name := c.Query("name")
    c.String(200, name)
  })

  h.Spin()
}
