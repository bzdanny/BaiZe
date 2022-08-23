package monitorController

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
)

type InfoServerController struct {
}

func NewInfoServerController() *InfoServerController {
	return &InfoServerController{}
}

func (isc *InfoServerController) GetInfoServer(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SuccessData(monitorModels.NewServer())
}
