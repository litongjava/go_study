package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitializeRedisClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6391",
		Password: "",
		DB:       0,
		// 连接池大小
		PoolSize: 100,
	})
	_, err = Rdb.Ping(Ctx).Result()
	return
}
