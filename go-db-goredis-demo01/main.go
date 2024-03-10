package main

import (
  "context"
  "fmt"
  "github.com/go-redis/redis/v8"
  appRedis "go-db-goredis-demo01/redis"
)

func main() {
  err := appRedis.InitializeRedisClient()
  if err != nil {
    panic(err)
  }

  val2, err := appRedis.Rdb.Get(appRedis.Ctx, "key2").Result()
  if err == redis.Nil {
    fmt.Println("key2 does not exist")
  } else if err != nil {
    panic(err)
  } else {
    fmt.Println("key2", val2)
  }

  hGetAllDemo(appRedis.Rdb, appRedis.Ctx)
}

func hGetAllDemo(rdb *redis.Client, ctx context.Context) {
  v := rdb.HGetAll(ctx, "user").Val()
  fmt.Println(v)
  v2 := rdb.HMGet(ctx, "user", "name", "age").Val()
  fmt.Println(v2)
  v3 := rdb.HGet(ctx, "user", "age")
  fmt.Println(v3)
}
