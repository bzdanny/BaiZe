package token

import (
	"github.com/bzdanny/BaiZe/app/constant/constants"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/setting"
	"github.com/bzdanny/BaiZe/baize/utils/redisUtils"
	"time"
)

var timeLive time.Duration

func Init() {
	timeLive = time.Duration(setting.Conf.TokenConfig.ExpireTime) * time.Minute
}

func RefreshToken(loginUser *systemModels.LoginUser) {
	loginUser.ExpireTime = time.Now().Add(timeLive).Unix()
	token := constants.LoginTokenKey + loginUser.Token
	redisUtils.SetStruct(token, loginUser, timeLive)
}
