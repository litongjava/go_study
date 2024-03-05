package handler

import (
  "context"
  "fmt"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/common/utils"
  "go-db-gorm-select-policy-demo01/container"
  "go-db-gorm-select-policy-demo01/models"
  "gorm.io/plugin/dbresolver"
  "time"
)

//sources笔replica快一半
func Time(ctx context.Context, c *app.RequestContext) {
  responseTime := getResponseTimeFromSources()
  responseTime2 := getResponseTimeFromReplica()
  c.JSON(200, utils.H{
    "responseTime":  responseTime,
    "responseTime2": responseTime2,
  })
}

func getResponseTimeFromReplica() int64 {
  var models []models.Feedback
  start := time.Now()
  container.Db.Find(&models)
  responseTime := time.Since(start).Milliseconds()
  return responseTime
}

func getResponseTimeFromSources() int64 {
  var models []models.Feedback
  start := time.Now()
  container.Db.Clauses(dbresolver.Write).Find(&models)
  responseTime := time.Since(start).Milliseconds()
  fmt.Println("responseTime", responseTime)
  return responseTime
}
