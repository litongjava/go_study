package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()

  //// 路由分组
  //r.GET("/posts", GetHandler)
  //r.POST("/posts", PostHandler)
  //// 删除id = 1 这篇文章
  //r.DELETE("/posts/1", DeleteHandler)

  //
  p := r.Group("/posts")
  {
    p.GET("", GetHandler)
    p.POST("", PostHandler)
    // 删除id = 1 这篇文章
    p.DELETE("/:id", DeleteHandler)
  }

  // localhost:8080/api/v1
  // Simple group: v1
  //v1 := r.Group("/api/v1")
  //{
  //	v1.POST("/login", loginEndpoint)
  //	v1.POST("/submit", submitEndpoint)
  //	v1.POST("/read", readEndpoint)
  //}

  // Simple group: v2
  //v2 := r.Group("/v2")
  //{
  //	v2.POST("/login", loginEndpoint)
  //	v2.POST("/submit", submitEndpoint)
  //	v2.POST("/read", readEndpoint)
  //}

  //
  api := r.Group("/api")
  {
    v1 := api.Group("v1")
    {
      v1.GET("posts", GetHandler)
      v1.POST("posts", PostHandler)
      // 删除id = 1 这篇文章
      v1.DELETE("posts/:id", DeleteHandler)
    }
  }

  r.Run()
}

func readEndpoint(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "readEndpoint",
  })
}

func submitEndpoint(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "submitEndpoint",
  })
}

func loginEndpoint(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "loginEndpoint",
  })
}

func DeleteHandler(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "DELETE",
  })
}

func PostHandler(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "POST",
  })
}

func GetHandler(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "GET",
  })
}
