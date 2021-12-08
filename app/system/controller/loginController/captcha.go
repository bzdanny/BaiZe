package loginController

import (
	"baize/app/common/commonModels"
	"baize/app/system/service/loginService/loginServiceImpl"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {
	code, err := loginServiceImpl.GenerateCode()
	if err != nil {
		c.JSON(http.StatusOK, commonModels.Error())
	} else {
		c.JSON(http.StatusOK, commonModels.SuccessData(code))
	}
}
