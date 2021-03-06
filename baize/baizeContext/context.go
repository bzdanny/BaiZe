package baizeContext

import (
	"github.com/bzdanny/BaiZe/app/constant/business"
	"github.com/bzdanny/BaiZe/app/constant/constants"
	monitorModels2 "github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/baize/utils/ipUtils"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
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
	ol := new(monitorModels2.SysOpenLog)
	ol.Title = title
	ol.BusinessType = businessTy.Msg()
	ol.Status = business.Success.Msg()
	ol.OperIp = bzc.ClientIP()
	ol.OperUrl = bzc.Request.URL.Path
	ol.RequestMethod = bzc.Request.Method
	loginUser := bzc.GetUser()
	if loginUser != nil {
		ol.OperName = loginUser.UserName
	}
	bzc.Set(constants.LogKey, ol)
}
func (bzc *BaiZeContext) SetUserAgent(login *monitorModels2.Logininfor) {
	login.InfoId = snowflake.GenID()
	ua := user_agent.New(bzc.Request.Header.Get("User-Agent"))
	ip := bzc.ClientIP()
	login.IpAddr = ip
	login.Os = ua.OS()
	login.LoginLocation = ipUtils.GetRealAddressByIP(ip)
	login.Browser, _ = ua.Browser()
}
