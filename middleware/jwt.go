package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"jin-gin/app/common/response"
	"jin-gin/app/services"
	"jin-gin/global"
)

// JWTAuth 校验token
func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(services.TokenType)+1:]

		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		// 判断token 是否正确 并且不在黑名单里
		if err != nil || services.JwtServices.IsInBlacklist(tokenStr) {
			response.TokenFail(c)
			c.Abort()
			return
		}
		claims := token.Claims.(*services.CustomClaims)
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Set("authId", claims.ID)
	}
}
