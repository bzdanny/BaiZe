package systemRouter

import (
	loginController2 "baize/app/system/controller/loginController"
	"github.com/gin-gonic/gin"
)

func InitGetUser(router *gin.RouterGroup) {
	router.GET("/getInfo", loginController2.GetInfo)
	router.GET("/getRouters", loginController2.GetRouters)
	router.POST("/logout", loginController2.Logout)
}
func InitLoginRouter(router *gin.RouterGroup) {
	router.GET("/captchaImage", loginController2.GetCode) //获取验证码
	router.POST("/login", loginController2.Login)         //登录
}
