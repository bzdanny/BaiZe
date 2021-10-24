package systemRouter

import (
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitGetUser(router *gin.RouterGroup) {
	router.GET("/getInfo", systemController.GetInfo)
	router.GET("/getRouters", systemController.GetRouters)
	router.POST("/logout", systemController.Logout)
}
func InitLoginRouter(router *gin.RouterGroup) {
	router.GET("/captchaImage", systemController.GetCode) //获取验证码
	router.POST("/login", systemController.Login)         //登录
}
