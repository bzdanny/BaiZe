package loginService

import (
	"baize/app/common/commonModels"
	"baize/app/system/models/loginModels"
	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	Login(login *loginModels.LoginBody, c *gin.Context) *commonModels.ResponseData
}
