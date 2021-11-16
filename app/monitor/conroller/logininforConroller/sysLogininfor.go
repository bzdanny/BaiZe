package logininforConroller

import (
	"baize/app/common/commonLog"
	"baize/app/common/commonModels"
	"baize/app/monitor/monitorModels"
	"baize/app/monitor/monitorService"
	"baize/app/monitor/monitorService/monitorServiceImpl"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var iLoginfor monitorService.ILogininforService = monitorServiceImpl.GetLogininforService()

func LogininforList(c *gin.Context) {
	loginfor := new(monitorModels.LogininforDQL)
	c.ShouldBind(loginfor)
	loginfor.SetLimit(c)
	list, count := iLoginfor.SelectLogininforList(loginfor)
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func LogininforExport(c *gin.Context) {
	commonLog.SetLog(c, "登录日志", "EXPORT")
}

func LogininforRemove(c *gin.Context) {
	commonLog.SetLog(c, "登录日志", "DELETE")
	var s slicesUtils.Slices = strings.Split(c.Param("infoIds"), ",")
	iLoginfor.DeleteLogininforByIds(s.StrSlicesToInt())
	c.JSON(http.StatusOK, commonModels.Success())

}

func LogininforClean(c *gin.Context) {
	commonLog.SetLog(c, "登录日志", "CLEAN")
	iLoginfor.CleanLogininfor()
	c.JSON(http.StatusOK, commonModels.Success())

}
