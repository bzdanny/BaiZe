package operLogController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/monitor/monitorModels"
	"baize/app/monitor/monitorService"
	"baize/app/monitor/monitorService/monitorServiceImpl"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"strings"
)

var iOperLog monitorService.ISysOperLogService = monitorServiceImpl.GetOperLogServiceService()

func OperLogList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	operLog := new(monitorModels.SysOpenLogDQL)
	c.ShouldBind(operLog)
	operLog.SetLimit(c)
	list, count := iOperLog.SelectOperLogList(operLog)
	bzc.SuccessListData(list, count)

}

func OperLogExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("操作日志", "EXPORT")
}

func OperLogRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("操作日志", "DELETE")
	var s slicesUtils.Slices = strings.Split(c.Param("operIds"), ",")
	iOperLog.DeleteOperLogByIds(s.StrSlicesToInt())
	bzc.Success()
}

func OperLogClean(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("操作日志", "CLEAN")
	iOperLog.CleanOperLog()
	bzc.Success()

}
