package logininforController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/monitor/monitorModels"
	"baize/app/monitor/monitorService"
	"baize/app/monitor/monitorService/monitorServiceImpl"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"strings"
)

var iLoginfor monitorService.ILogininforService = monitorServiceImpl.GetLogininforService()

func LogininforList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginfor := new(monitorModels.LogininforDQL)
	c.ShouldBind(loginfor)
	loginfor.SetLimit(c)
	list, count := iLoginfor.SelectLogininforList(loginfor)
	bzc.SuccessListData(list, count)

}

func LogininforExport(c *gin.Context) {

}

func LogininforRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("登录日志", "DELETE")
	var s slicesUtils.Slices = strings.Split(c.Param("infoIds"), ",")
	iLoginfor.DeleteLogininforByIds(s.StrSlicesToInt())
	bzc.Success()
}

func LogininforClean(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("登录日志", "CLEAN")
	iLoginfor.CleanLogininfor()
	bzc.Success()

}
