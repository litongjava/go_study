package main

import "github.com/gin-gonic/gin"

type User struct {
  // application/json json
  // application/x-www-form-urlencoded form
  ID string `form:"id" binding:"required"`
}

func main() {
  r := gin.Default()
  r.POST("user", func(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
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
  })
  r.Run()
}
