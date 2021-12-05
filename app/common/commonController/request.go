package commonController

import (
	"baize/app/constant/constants"
	"baize/app/system/models/loginModels"
	"github.com/gin-gonic/gin"
)

func GetCurrentLoginUser(c *gin.Context) (loginUser *loginModels.LoginUser) {
	loginUserKey, _ := c.Get(constants.LoginUserKey)
	if loginUserKey != nil {
		loginUser = loginUserKey.(*loginModels.LoginUser)
	}

	return
}
