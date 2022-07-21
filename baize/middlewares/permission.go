package middlewares

import (
	"github.com/bzdanny/BaiZe/app/constant/constants"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/utils"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/bzdanny/BaiZe/baize/baizeSet"
	"github.com/gin-gonic/gin"
)

func HasPermissionSet(p baizeSet.Set[string]) func(c *gin.Context) {
	return func(c *gin.Context) {
		bzc := baizeContext.NewBaiZeContext(c)

		loginUserKey, _ := c.Get(constants.LoginUserKey)
		user := loginUserKey.(*systemModels.LoginUser)
		if utils.IsAdmin(user.User.UserId) {
			c.Next()
			return
		}
		permissions := user.Permissions
		if permissions == nil || len(permissions) == 0 {
			bzc.PermissionDenied()
			c.Abort()
			return
		}
		if !hasPermissionsSet(permissions, p) {
			bzc.PermissionDenied()
			c.Abort()
			return
		}
		c.Next()
	}
}

func hasPermissionsSet(permissions []string, permission baizeSet.Set[string]) bool {
	for _, v := range permissions {
		if permission.Contains(v) {
			return true
		}
	}
	return false
}

func HasPermission(permission string) func(c *gin.Context) {
	return func(c *gin.Context) {
		bzc := baizeContext.NewBaiZeContext(c)
		if permission == "" {
			c.Next()
			return
		}
		loginUserKey, _ := c.Get(constants.LoginUserKey)
		user := loginUserKey.(*systemModels.LoginUser)
		if utils.IsAdmin(user.User.UserId) {
			c.Next()
			return
		}
		permissions := user.Permissions
		if permissions == nil || len(permissions) == 0 {
			bzc.PermissionDenied()
			c.Abort()
			return
		}
		if !hasPermissions(permissions, permission) {
			bzc.PermissionDenied()
			c.Abort()
			return
		}
		c.Next()
	}
}

func hasPermissions(permissions []string, permission string) bool {
	for _, v := range permissions {
		if v == permission {
			return true
		}
	}
	return false
}
