package monitorController

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

var iOperLog monitorService.ISysOperLogService = monitorServiceImpl.GetOperLogServiceService()

func OperLogList(c *gin.Context) {
	operLog := new(monitorModels.SysOpenLogDQL)
	c.ShouldBind(operLog)
	operLog.SetLimit(c)
	list, count := iOperLog.SelectOperLogList(operLog)
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func OperLogExport(c *gin.Context) {
	commonLog.SetLog(c, "操作日志", "EXPORT")
}

func OperLogRemove(c *gin.Context) {
	commonLog.SetLog(c, "操作日志", "DELETE")
	var s slicesUtils.Slices = strings.Split(c.Param("operIds"), ",")
	iOperLog.DeleteOperLogByIds(s.StrSlicesToInt())
	c.JSON(http.StatusOK, commonModels.Success())
}

func OperLogClean(c *gin.Context) {
	commonLog.SetLog(c, "操作日志", "CLEAN")
	iOperLog.CleanOperLog()
	c.JSON(http.StatusOK, commonModels.Success())

}
