package main

import "github.com/gin-gonic/gin"

type User struct {
  ID string `uri:"id" binding:"required,uuid"`
}

func main() {
  r := gin.Default()
  // http://localhost:8080/user/1234
  r.GET("user/:id", func(c *gin.Context) {
    var user User
    if err := c.ShouldBindUri(&user); err != nil {
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
