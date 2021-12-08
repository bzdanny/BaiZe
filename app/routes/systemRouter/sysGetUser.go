package systemRouter

import (
	"baize/app/system/controller/loginController"
	"github.com/gin-gonic/gin"
)

func InitGetUser(router *gin.RouterGroup) {
	router.GET("/getInfo", loginController.GetInfo)
	router.GET("/getRouters", loginController.GetRouters)

}
func InitLoginRouter(router *gin.RouterGroup) {
	router.GET("/captchaImage", loginController.GetCode) //获取验证码
	router.POST("/login", loginController.Login)         //登录
	router.POST("/logout", loginController.Logout)
}
