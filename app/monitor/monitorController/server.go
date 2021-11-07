package monitorController

import (
	"baize/app/common/commonModels"
	"baize/app/monitor/monitorModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInfoServer(c *gin.Context) {
	c.JSON(http.StatusOK, commonModels.SuccessData(monitorModels.NewServer()))
}
