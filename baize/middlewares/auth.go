package middlewares

import (
	"github.com/bzdanny/BaiZe/app/constant/constants"
	"github.com/bzdanny/BaiZe/app/utils/jwt"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/bzdanny/BaiZe/baize/utils/token"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		bzc := baizeContext.NewBaiZeContext(c)
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			bzc.InvalidToken()
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == constants.TokenPrefix) {
			bzc.InvalidToken()
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		loginUser, err := jwt.ParseToken(parts[1])
		if err != nil {
			bzc.InvalidToken()
			c.Abort()
			return
		}
		if loginUser.ExpireTime < time.Now().Add(time.Duration(15)*time.Minute).Unix() {
			go token.RefreshToken(loginUser)
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(constants.LoginUserKey, loginUser)

		c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
	}
}
