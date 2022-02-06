package serverController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/monitor/monitorModels"
	"github.com/gin-gonic/gin"
)

func GetInfoServer(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SuccessData(monitorModels.NewServer())
}
