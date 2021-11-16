package systemRouter

import (
	"baize/app/system/controller/ProfileController"
	"github.com/gin-gonic/gin"
)

func InitSysProfileRouter(router *gin.RouterGroup) {
	systemProfile := router.Group("/system/user/profile")

	systemProfile.GET("", ProfileController.Profile)
	systemProfile.PUT("", ProfileController.ProfileUpdateProfile)
	systemProfile.PUT("/updatePwd", ProfileController.ProfileUpdatePwd)
	systemProfile.POST("/avatar", ProfileController.ProfileAvatar)
}
