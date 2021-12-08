package utils

import (
	setting "baize/app/setting"
	jwt2 "baize/app/utils/jwt"
	snowflake "baize/app/utils/snowflake"
	token "baize/app/utils/token"
)

func Init() {
	jwt2.Init(setting.Conf.TokenConfig)
	token.Init(setting.Conf.TokenConfig.ExpireTime)
	snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID)

}
