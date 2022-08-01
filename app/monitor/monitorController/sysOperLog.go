package monitorController

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/app/monitor/monitorService"
	"github.com/bzdanny/BaiZe/app/monitor/monitorService/monitorServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
)

type OperLogController struct {
	os monitorService.ISysOperLogService
}

func NewOperLogController(os *monitorServiceImpl.OperLogService) *OperLogController {
	return &OperLogController{
		os: os,
	}
}

func (olc *OperLogController) OperLogList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	operLog := new(monitorModels.SysOpenLogDQL)
	c.ShouldBind(operLog)
	list, count := olc.os.SelectOperLogList(operLog)
	bzc.SuccessListData(list, count)

}

func (olc *OperLogController) OperLogExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("操作日志", "EXPORT")
}

func (olc *OperLogController) OperLogRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("操作日志", "DELETE")
	olc.os.DeleteOperLogByIds(bzc.ParamInt64Array("operIds"))
	bzc.Success()
}

func (olc *OperLogController) OperLogClean(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("操作日志", "CLEAN")
	olc.os.CleanOperLog()
	bzc.Success()

}
