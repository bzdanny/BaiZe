package redisUtils

import (
	"encoding/json"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"

	"time"
)

func SetString(key string, str string, expiration time.Duration) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("redisUtils", zap.Any("error", err))
			}
		}()
		err := datasource.GetRedisClient().Set(key, str, expiration).Err()
		if err != nil {
			zap.L().Error("Redis存储失败", zap.Error(err))
		}
	}()
}

func GetString(key string) string {
	return datasource.GetRedisClient().Get(key).Val()
}
func GetInt64(key string) int64 {
	return gconv.Int64(GetString(key))
}

func SetStruct(key string, value interface{}, expiration time.Duration) {
	marshal, err := json.Marshal(value)
	if err != nil {
		zap.L().Error("Redis存储失败", zap.Error(err))
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("redisUtils", zap.Any("error", err))
			}
		}()
		err = datasource.GetRedisClient().Set(key, marshal, expiration).Err()
		if err != nil {
			zap.L().Error("Redis存储失败", zap.Error(err))
		}
	}()
}

func Keys(pattern string) []string {
	cmd := datasource.GetRedisClient().Keys(pattern)
	return cmd.Val()
}

func Delete(key string) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("redisUtils", zap.Any("error", err))
			}
		}()
		datasource.GetRedisClient().Del(key)
	}()
}

func GetStruct[T any](key string, tt *T) (*T, error) {
	newT := new(T)
	LoginUserJson := GetString(key)
	err := json.Unmarshal([]byte(LoginUserJson), newT)
	return newT, err
}
