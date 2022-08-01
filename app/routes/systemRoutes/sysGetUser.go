package systemRoutes

import (
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/gin-gonic/gin"
)

func InitGetUser(router *gin.RouterGroup, loginController *systemController.LoginController) {
	router.GET("/getInfo", loginController.GetInfo)
	router.GET("/getRouters", loginController.GetRouters)

}
func InitLoginRouter(router *gin.RouterGroup, loginController *systemController.LoginController) {
	router.GET("/captchaImage", loginController.GetCode) //获取验证码
	router.POST("/login", loginController.Login)         //登录
	router.POST("/logout", loginController.Logout)
}
