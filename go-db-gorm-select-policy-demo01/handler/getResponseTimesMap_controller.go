package handler

import (
  "context"
  "fmt"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/common/utils"
  "go-db-gorm-select-policy-demo01/gorm/policy"
)

func GetResponseTimesMap(ctx context.Context, c *app.RequestContext) {
  policy.ResponseTimesMap.Range(func(key, value interface{}) bool {
    fmt.Println("Key:", key, "Value:", value)
    return true // 继续遍历
  })
  c.JSON(200, utils.H{"data": policy.ResponseTimesMap})

}
