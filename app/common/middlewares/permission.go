package middlewares

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/constant/constants"
	"baize/app/system/models/loginModels"
	"baize/app/utils/admin"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
)

func HasPermission(permission string) func(c *gin.Context) {
	return func(c *gin.Context) {
		bzc := baizeContext.NewBaiZeContext(c)
		if permission == "" {
			c.Next()
			return
		}
		loginUserKey, _ := c.Get(constants.LoginUserKey)
		user := loginUserKey.(*loginModels.LoginUser)
		if admin.IsAdmin(user.User.UserId) {
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
	var s slicesUtils.Slices = permissions
	return s.Contains(permission) > -1

}
