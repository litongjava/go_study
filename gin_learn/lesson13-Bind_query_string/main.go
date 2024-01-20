package main

import "github.com/gin-gonic/gin"

type User struct {
  ID       string `form:"id" binding:"required,uuid"`
  UserName string `form:"username" binding:"required,min=3"`
  PassWord string `form:"password" binding:"required,min=3"`
}

func main() {
  r := gin.Default()
  r.POST("user", userHandler())
  r.Run()
}

func userHandler() func(c *gin.Context) {
  return func(c *gin.Context) {
    var user User
    if err := c.ShouldBindQuery(&user); err != nil {
      c.JSON(200, gin.H{
        "Code": 200,
        "Msg":  err.Error(),
      })
      return
    }
    c.JSON(200, gin.H{
      "Code": 0,
      "ID":   user.ID,
    })
  }
}
