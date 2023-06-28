package cache

import (
	"fmt"
	"gin_mall/conf"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	InitRedisClient()
}

func InitRedisClient() {
	client := redis.NewClient(&redis.Options{
		Addr: conf.RedisConf.RedisAddr,
		DB:   conf.RedisConf.RedisDb,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}

// product key
func ProductViewKey(pid uint) string {
	return fmt.Sprintf("product:view:%d", pid)
}
