package systemController

import (
	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	code := loginServiceImpl.GenerateCode()
	bzc.SuccessData(code)

}
