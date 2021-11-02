package loginService

import (
	"baize/app/monitor/monitorModels"
	"baize/app/utils/IpUtils"
	"baize/app/utils/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func setUserAgent(login *monitorModels.Logininfor, c *gin.Context) {
	login.InfoId = snowflake.GenID()
	ua := user_agent.New(c.Request.Header.Get("User-Agent"))
	ip := c.ClientIP()
	login.IpAddr = ip
	login.Os = ua.OS()
	login.LoginLocation = IpUtils.GetRealAddressByIP(ip)
	login.Browser, _ = ua.Browser()

}
