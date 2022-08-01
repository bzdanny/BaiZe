package monitorController

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorService"
	"github.com/bzdanny/BaiZe/app/monitor/monitorService/monitorServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
)

type UserOnlineController struct {
	uos monitorService.IUserOnlineService
}

func NewUserOnlineController(uos *monitorServiceImpl.UserOnlineService) *UserOnlineController {
	return &UserOnlineController{
		uos: uos,
	}
}

func (uoc *UserOnlineController) UserOnlineList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	list, total := uoc.uos.SelectUserOnlineList(c.Query("ipaddr"), c.Query("userName"))
	bzc.SuccessListData(list, total)
}
func (uoc *UserOnlineController) ForceLogout(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("在线用户", "FORCE")
	uoc.uos.ForceLogout(c.Param("tokenId"))
	bzc.Success()
}
