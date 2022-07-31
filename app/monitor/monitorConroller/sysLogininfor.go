package monitorConroller

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/app/monitor/monitorService"
	"github.com/bzdanny/BaiZe/app/monitor/monitorService/monitorServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
)

type LogininforController struct {
	ls monitorService.ILogininforService
}

func NewLogininforController(ls *monitorServiceImpl.LogininforService) *LogininforController {
	return &LogininforController{ls: ls}
}

func (lc *LogininforController) LogininforList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginfor := new(monitorModels.LogininforDQL)
	c.ShouldBind(loginfor)
	list, count := lc.ls.SelectLogininforList(loginfor)
	bzc.SuccessListData(list, count)

}

func (lc *LogininforController) LogininforExport(c *gin.Context) {

}

func (lc *LogininforController) LogininforRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("登录日志", "DELETE")
	lc.ls.DeleteLogininforByIds(bzc.ParamInt64Array("infoIds"))
	bzc.Success()
}

func (lc *LogininforController) LogininforClean(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("登录日志", "CLEAN")
	lc.ls.CleanLogininfor()
	bzc.Success()

}
