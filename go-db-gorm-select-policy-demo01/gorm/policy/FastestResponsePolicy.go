package policy

import (
  "errors"
  "github.com/cloudwego/hertz/pkg/common/hlog"
  "gorm.io/gorm"
  "log"
  "sync"
)

// ResponseTimeInfo 用于存储累加的响应时间和查询次数
type ResponseTimeInfo struct {
  TotalResponseTime int64 // 累计响应时间，以毫秒为单位
  QueryCount        int64 // 查询次数
  mu                sync.Mutex
}

var ResponseTimesMap sync.Map // 存储每个 ConnPool 的 ResponseTimeInfo

// 更新响应时间的函数，现在会计算平均值
func UpdateResponseTime(poolIdentifier gorm.ConnPool, responseTimeMillis int64) {
  // 从 map 中获取当前池的 ResponseTimeInfo
  v, _ := ResponseTimesMap.LoadOrStore(poolIdentifier, &ResponseTimeInfo{})
  info := v.(*ResponseTimeInfo)

  // 锁定以更新信息
  info.mu.Lock()
  defer info.mu.Unlock()

  info.TotalResponseTime += responseTimeMillis
  info.QueryCount++
}

// 获取平均响应时间
func GetAverageResponseTime(poolIdentifier gorm.ConnPool) (avgRT float64, ok bool) {
  v, ok := ResponseTimesMap.Load(poolIdentifier)
  if !ok {
    return 0, false
  }
  info := v.(*ResponseTimeInfo)

  info.mu.Lock()
  defer info.mu.Unlock()

  if info.QueryCount == 0 {
    return 0, true // 避免除零错误
  }
  avgRT = float64(info.TotalResponseTime) / float64(info.QueryCount)
  return avgRT, true
}

type FastestResponsePolicy struct{}

// Resolve 方法根据记录的响应时间选择最快的连接池
func (p *FastestResponsePolicy) Resolve(pools []gorm.ConnPool) gorm.ConnPool {
  if len(pools) == 0 {
    return nil // 如果没有连接池，则直接返回nil
  }
  err := errors.New("error")
  if err != nil {
    log.Printf("Error occurred: %+v\n", err)
  }

  for _, pool := range pools {
    hlog.Infof("address: %p \n", &pool)
  }
  minAvgRT := float64(^uint(0) >> 1) // 初始化为最大float64值
  var selectedPool gorm.ConnPool
  found := false

  var elementIndex = 0
  for index, pool := range pools {
    if avgRT, ok := GetAverageResponseTime(pool); ok {
      // 如果找到平均响应时间且比当前最小值要小，则更新选中的连接池
      if avgRT < minAvgRT {
        selectedPool = pool
        minAvgRT = avgRT
        found = true
        elementIndex = index
      }
    }
  }

  hlog.Info("connect pool index:", elementIndex)
  if !found {
    return pools[elementIndex]
  }

  return selectedPool
}
