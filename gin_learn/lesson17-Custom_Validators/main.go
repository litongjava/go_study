package main

import (
  "strings"

  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
  "github.com/go-playground/validator/v10"
)

type User struct {
  // 需求：id如果非1开头则验证失败，反之
  Id string `form:"id" binding:"required,micheal"`
}

var customValidate validator.Func = func(fl validator.FieldLevel) bool {
  data := fl.Field().Interface().(string)
  if strings.HasPrefix(data, "1") {
    return true
  }
  return false
}

func main() {
  r := gin.Default()
  // 注册
  if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    _ = v.RegisterValidation("micheal", customValidate)
  }
  r.GET("user", func(c *gin.Context) {
    var u User
    if err := c.ShouldBind(&u); err != nil {
      c.JSON(200, gin.H{
        "Code": 200,
        "Msg":  err.Error(),
      })
      return
    }
    c.JSON(200, gin.H{
      "Code": 200,
      "Msg":  u.Id,
    })
  })
  r.Run()
}
