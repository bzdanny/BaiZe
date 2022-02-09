package baizeContext

import (
	"baize/app/constant/business"
	"baize/app/constant/constants"
	"baize/app/monitor/monitorModels"
	"baize/app/utils/IpUtils"
	"baize/app/utils/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

type BaiZeContext struct {
	*gin.Context
}

func NewBaiZeContext(c *gin.Context) *BaiZeContext {
	return &BaiZeContext{c}
}

func (bzc *BaiZeContext) SetLog(title string, businessTy business.BusinessType) {
	ol := new(monitorModels.SysOpenLog)
	ol.Title = title
	ol.BusinessType = businessTy.Msg()
	ol.Status = business.Success.Msg()
	ol.OperIp = bzc.ClientIP()
	ol.OperUrl = bzc.Request.URL.Path
	ol.RequestMethod = bzc.Request.Method
	loginUser := bzc.GetCurrentLoginUser()
	if loginUser != nil {
		ol.OperName = loginUser.User.UserName
	}
	bzc.Set(constants.LogKey, ol)
}
func (bzc *BaiZeContext) SetUserAgent(login *monitorModels.Logininfor) {
	login.InfoId = snowflake.GenID()
	ua := user_agent.New(bzc.Request.Header.Get("User-Agent"))
	ip := bzc.ClientIP()
	login.IpAddr = ip
	login.Os = ua.OS()
	login.LoginLocation = IpUtils.GetRealAddressByIP(ip)
	login.Browser, _ = ua.Browser()
}
