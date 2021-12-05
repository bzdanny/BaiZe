package redis

import (
	"baize/app/system/models/loginModels"
	"encoding/json"
	"time"

	"go.uber.org/zap"
)

func SetString(key string, str string, expiration time.Duration) {
	err := client.Set(key, str, expiration).Err()
	if err != nil {
		zap.L().Error("redis存储错误", zap.Error(err))
	}
}

func SetStruct(key string, value interface{}, expiration time.Duration) {
	marshal, err := json.Marshal(value)
	if err != nil {
		zap.L().Error("json转换错误", zap.Error(err))
	}
	err = client.Set(key, marshal, expiration).Err()
	if err != nil {
		zap.L().Error("redis存储错误", zap.Error(err))
	}
}

func GetString(key string) (val string) {
	val = client.Get(key).Val()
	return
}

func GetCacheLoginUser(key string) (*loginModels.LoginUser, error) {
	loginUser := new(loginModels.LoginUser)
	LoginUserJson := GetString(key)
	err := json.Unmarshal([]byte(LoginUserJson), loginUser)
	return loginUser, err
}
func Keys(pattern string) []string {
	client.Del()
	cmd := client.Keys(pattern)
	return cmd.Val()
}

func Delete(key string) {
	client.Del(key)
}
