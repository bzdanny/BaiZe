package loginController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/system/service/loginService/loginServiceImpl"
	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	code := loginServiceImpl.GenerateCode()
	bzc.SuccessData(code)

}
