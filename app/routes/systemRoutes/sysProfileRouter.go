package systemRoutes

import (
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysProfileRouter(router *gin.RouterGroup, profileController *systemController.ProfileController) {
	systemProfile := router.Group("/system/user/profile")

	systemProfile.GET("", profileController.Profile)
	systemProfile.PUT("", profileController.ProfileUpdateProfile)
	systemProfile.PUT("/updatePwd", profileController.ProfileUpdatePwd)
	systemProfile.POST("/avatar", profileController.ProfileAvatar)
}
