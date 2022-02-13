package systemRouter

import (
	"baize/app/system/controller/profileController"
	"github.com/gin-gonic/gin"
)

func InitSysProfileRouter(router *gin.RouterGroup) {
	systemProfile := router.Group("/system/user/profile")

	systemProfile.GET("", profileController.Profile)
	systemProfile.PUT("", profileController.ProfileUpdateProfile)
	systemProfile.PUT("/updatePwd", profileController.ProfileUpdatePwd)
	systemProfile.POST("/avatar", profileController.ProfileAvatar)
}
