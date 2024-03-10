package redis

import (
  "context"
  "github.com/go-redis/redis/v8"
  "time"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitializeRedisClient() (err error) {
  opt, err := redis.ParseURL("redis://root:@localhost:6379/0")
  if err != nil {
    panic(err)
  }
  Rdb = redis.NewClient(opt)
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  _, err = Rdb.Ping(ctx).Result()
  return err
}

//func InitializeRedisClient() (err error) {
//  Rdb = redis.NewClient(&redis.Options{
//    Addr:     "localhost:6391",
//    Password: "",
//    DB:       0,
//    // 连接池大小
//    PoolSize: 100,
//  })
//  _, err = Rdb.Ping(Ctx).Result()
//  return
//}

//func initializeRedisClient() error {
//  Rdb = redis.NewFailoverClient(&redis.FailoverOptions{
//    MasterName:    "master",
//    SentinelAddrs: []string{"10.0.0.10:6379", "10.0.0.11:6379", "10.0.0.11:6379"},
//  })
//  _, err := Rdb.Ping(Ctx).Result()
//  return err
//}

func initializeRedisClient() error {
  //var client *redis.ClusterClient
  client := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7006"},
  })
  _, err := client.Ping(Ctx).Result()
  return err
}
