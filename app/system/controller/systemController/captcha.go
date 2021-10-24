package systemController

import (
	"baize/app/common/commonModels"
	"baize/app/system/service/loginService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {
	code, err := loginService.GenerateCode()
	if err != nil {
		c.JSON(http.StatusOK, commonModels.Error())
	} else {
		c.JSON(http.StatusOK, commonModels.SuccessData(code))
	}
}
