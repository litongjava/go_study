package main

import "github.com/gin-gonic/gin"

type Person struct {
  ID   int    `uri:"id"`
  Name string `uri:"name"`
}

func main() {
  r := gin.Default()

  //r.GET("/posts", func(c *gin.Context) {
  //	c.String(200, "1")
  //})

  //r.GET("/:id/:name", func(c *gin.Context) {
  //	id := c.Param("id")
  //	name := c.Param("name")
  //	c.JSON(200, gin.H{
  //		"id":     id,
  //		"name": name,
  //	})
  //})

  r.GET("/:id/:name", func(c *gin.Context) {
    var p Person
    if err := c.ShouldBindUri(&p); err != nil {
      c.Status(404)
      return
    }
    c.JSON(200, gin.H{
      "name": p.Name,
      "id":   p.ID,
    })
  })

  // 删除id=1这条数据
  //r.DELETE("posts/:id", func(c *gin.Context) {
  //
  //})

  r.Run()
}
