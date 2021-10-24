package loginService

import (
	"baize/app/system/models/systemModels"
	"baize/app/utils/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func setUserAgent(login *systemModels.Logininfor, c *gin.Context) {
	login.InfoId = snowflake.GenID()
	ua := user_agent.New(c.Request.Header.Get("User-Agent"))
	login.IpAddr = c.ClientIP()
	login.Os = ua.OS()
	//TODO   login.LoginLocation
	login.Browser, _ = ua.Browser()

}
