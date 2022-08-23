package baizeContext

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/baize/utils/ipUtils"
	"github.com/bzdanny/BaiZe/baize/utils/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

type BaiZeContext struct {
	*gin.Context
}

func NewBaiZeContext(c *gin.Context) *BaiZeContext {
	return &BaiZeContext{c}
}

func (bzc *BaiZeContext) SetUserAgent(login *monitorModels.Logininfor) {
	login.InfoId = snowflake.GenID()
	ua := user_agent.New(bzc.Request.Header.Get("User-Agent"))
	ip := bzc.ClientIP()
	login.IpAddr = ip
	login.Os = ua.OS()
	login.LoginLocation = ipUtils.GetRealAddressByIP(ip)
	login.Browser, _ = ua.Browser()
}
