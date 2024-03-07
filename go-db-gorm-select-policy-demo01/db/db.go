package db

import (
  "github.com/cloudwego/hertz/pkg/common/hlog"
  "go-db-gorm-select-policy-demo01/config"
  "go-db-gorm-select-policy-demo01/container"
  "go-db-gorm-select-policy-demo01/gorm/policy"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/plugin/dbresolver"
  "reflect"
  "strings"
  "time"
)

func Init() {
  //gorm.Dialector
  dialector := postgres.Open(config.DATABASE_DSN)
  db, err := gorm.Open(dialector, &gorm.Config{
    SkipDefaultTransaction: true,
    PrepareStmt:            true,
  })
  if err != nil {
    panic(err)
  }

  registerCallBack(db)
  err = loadReplica(db)
  if err != nil {
    panic(err)
  }

  container.Db = db
}

func registerCallBack(db *gorm.DB) {
  db.Callback().Query().Register("update_response_time:before", func(db *gorm.DB) {
    db.InstanceSet("start_time", time.Now().UnixMilli())
  })

  db.Callback().Query().Register("update_response_time:after", func(db *gorm.DB) {
    // 在查询结束后执行
    startTime, ok := db.InstanceGet("start_time")
    if !ok {
      hlog.Error("Failed to get start time from DB instance")
      return
    }
    // 确保类型断言成功
    startMilli, ok := startTime.(int64)
    if !ok {
      hlog.Error("Start time is not of type int64")
      return
    }
    // 计算持续时间
    endMilli := time.Now().UnixMilli()
    elapsedTime := endMilli - startMilli
    hlog.Info("Query execution time (ms): ", endMilli, startMilli, elapsedTime)

    hlog.Info("address:", reflect.ValueOf(db.Statement.ConnPool).Pointer())
    hlog.Info("address:", reflect.ValueOf(db.ConnPool).Pointer())

    // 更新响应时间统计
    policy.UpdateResponseTime(db.Statement.ConnPool, elapsedTime)
  })
}

func loadReplica(db *gorm.DB) error {
  sources := make([]gorm.Dialector, 0)
  sourcesStr := strings.Split(config.DATABASE_DSN, ",")
  for _, dsn := range sourcesStr {
    dialector := postgres.Open(dsn)
    sources = append(sources, dialector)
    conn, err := gorm.Open(dialector, &gorm.Config{})
    if err != nil {
      return err
    }
    container.PoolToDSNMap[conn.ConnPool] = dsn
  }
  replicas := make([]gorm.Dialector, 0)
  replicasStr := strings.Split(config.DATABASE_REPLICAS, ",")
  for _, dsn := range replicasStr {
    dialector := postgres.Open(dsn)
    replicas = append(replicas, dialector)

    conn, err := gorm.Open(dialector, &gorm.Config{})
    if err != nil {
      return err
    }
    container.PoolToDSNMap[conn.ConnPool] = dsn
  }
  // 使用服务器DB实例上的`Use`方法注册新的解析器。
  // 此调用配置了数据库连接，用于在多个数据库间分配读/写操作。
  config := dbresolver.Config{
    Sources:  sources,                         // 指定主要/源数据库。
    Replicas: replicas,                        // 指定副本数据库。
    Policy:   &policy.FastestResponsePolicy{}, // 定义选择副本的策略；在这里是随机选择。
    //Policy:            dbresolver.RandomPolicy{},
    TraceResolverMode: true, // 开启解析器的跟踪模式。
  }

  dbResolver := dbresolver.Register(config)
  err := db.Use(dbResolver)
  return err
}
