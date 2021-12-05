package redis

import (
	"baize/app/setting"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化连接
func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", setting.Conf.RedisConfig.Host, setting.Conf.RedisConfig.Port),
		Password:     setting.Conf.RedisConfig.Password, // no password set
		DB:           setting.Conf.RedisConfig.DB,       // use default DB
		PoolSize:     setting.Conf.RedisConfig.PoolSize,
		MinIdleConns: setting.Conf.RedisConfig.MinIdleConns,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func Close() {
	_ = client.Close()
}
