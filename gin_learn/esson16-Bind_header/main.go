package main

import "github.com/gin-gonic/gin"

type Header struct {
  Referer string `header:"Referer" binding:"required"`
}

func main() {
  r := gin.Default()
  r.GET("user", func(c *gin.Context) {
    var h Header
    if err := c.ShouldBindHeader(&h); err != nil {
      c.JSON(200, gin.H{
        "Code": 200,
        "Msg":  err.Error(),
      })
      return
    }
    c.JSON(200, gin.H{
      "Code":    0,
      "Referer": h.Referer,
    })
  })
  r.Run()
}
