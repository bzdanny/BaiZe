package systemRouter

import (
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysProfileRouter(router *gin.RouterGroup) {
	systemProfile := router.Group("/system/user/profile")

	systemProfile.GET("", systemController.Profile)
	systemProfile.PUT("", systemController.ProfileUpdateProfile)
	systemProfile.PUT("/updatePwd", systemController.ProfileUpdatePwd)
	systemProfile.POST("/avatar", systemController.ProfileAvatar)
}
