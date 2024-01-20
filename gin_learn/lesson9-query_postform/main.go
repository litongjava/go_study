package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()
  r.POST("/user/:id", func(c *gin.Context) {
    id := c.Param("id")
    userName := c.PostForm("user_name")
    c.JSON(200, gin.H{
      "id":        id,
      "user_name": userName,
    })
  })
  r.Run()
}
