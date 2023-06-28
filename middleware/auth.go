package middleware

import (
	"gin_mall/pkg/util"
	"gin_mall/response"
	"time"

	"github.com/gin-gonic/gin"
)

// 用户登录认证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := response.SuccessCode
		msg := response.Success
		var claims *util.Claims

		token := c.GetHeader("Authorization")
		if token == "" {
			code = response.ErrorAuthCode
			msg = response.ErrorNotLogin
		} else {
			var err error
			claims, err = util.ParseToken(token)
			if err != nil {
				code = response.ErrorAuthCode
				msg = response.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt { // token 过期
				code = response.ErrorAuthCode
				msg = response.ErrorAuthCheckTokenTimeout
			}
		}

		if code != response.SuccessCode {
			response.FailResponse(c, code, msg)
			c.Abort()
			return
		}

		c.Set("ID", claims.ID)
		c.Set("UserName", claims.UserName)
		c.Set("Authority", claims.Authority)
		c.Next()
	}
}
