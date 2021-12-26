package swaggerTest

import (
	"baize/app/swaggerTest/testController"
	"github.com/gin-gonic/gin"
)

func InitGenTableRouter(router *gin.RouterGroup) {
	testUser := router.Group("/test/user")
	testUser.GET("/list",  testController.DemoUserList)
	testUser.GET(":userId", testController.GetUser)
	testUser.POST("", testController.Save)
	testUser.PUT("", testController.Update)
	testUser.DELETE(":userId", testController.Delete)
}