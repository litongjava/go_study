package handler

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/common/utils"
  "go-db-gorm-select-policy-demo01/container"
  "go-db-gorm-select-policy-demo01/models"
)

func FeedbackList(ctx context.Context, c *app.RequestContext) {
  var entities []models.Feedback
  container.Db.Find(&entities)
  c.JSON(200, utils.H{"data": entities})
}
