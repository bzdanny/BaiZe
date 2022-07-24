package token

import (
	"github.com/bzdanny/BaiZe/app/constant/constants"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/utils/redisUtils"
	"time"
)

var timeLive time.Duration

func Init(expireTime int64) {
	timeLive = time.Duration(expireTime) * time.Minute
}

func RefreshToken(loginUser *systemModels.LoginUser) {
	loginUser.ExpireTime = time.Now().Add(timeLive).Unix()
	token := constants.LoginTokenKey + loginUser.Token
	redisUtils.SetStruct(token, loginUser, timeLive)

}
