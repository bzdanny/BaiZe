package token

import (
	"baize/app/common/redis"
	"baize/app/constant/constants"
	"baize/app/system/models/loginModels"
	"time"
)

var timeLive time.Duration

func Init(expireTime int64) {
	timeLive = time.Duration(expireTime) * time.Minute
}

func RefreshToken(loginUser *loginModels.LoginUser) {
	loginUser.ExpireTime = time.Now().Add(timeLive).Unix()
	token := constants.LoginTokenKey + loginUser.Token
	redis.SetStruct(token, loginUser, timeLive)

}
