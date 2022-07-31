package monitorConroller

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

func GetInfoServer(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SuccessData(monitorModels.NewServer())
}
