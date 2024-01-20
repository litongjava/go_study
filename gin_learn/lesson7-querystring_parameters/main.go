package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()
  r.GET("welcome", welcomeHandler)
  r.GET("array", arrayHandler)
  r.GET("map", mapHandler)
  r.Run()
}

func mapHandler(c *gin.Context) {
  m := c.QueryMap("user")
  c.JSON(200, gin.H{
    "data": m,
  })
}

func arrayHandler(c *gin.Context) {
  ids := c.QueryArray("ids")
  c.JSON(200, gin.H{
    "ids": ids,
  })
}

func welcomeHandler(c *gin.Context) {
  // firstname lastname
  firstName := c.DefaultQuery("firstname", "micheal")
  lastName := c.Query("lastname")
  c.JSON(200, gin.H{
    "firstname": firstName,
    "lastname":  lastName,
  })

}
