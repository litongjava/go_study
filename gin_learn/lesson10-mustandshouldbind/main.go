package main

import "github.com/gin-gonic/gin"

type User struct {
  // edc24790-ad34-4666-a9d5-243cdb94071d
  ID string `form:"id" binding:"required,uuid"`
}

func main() {
  r := gin.Default()
  r.GET("user/", userHandler)
  r.Run()
}
func userHandler(c *gin.Context) {
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
    "Id":   user.ID,
  })
}
