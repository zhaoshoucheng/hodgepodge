package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoshoucheng/hodgepodge/Access"
)


func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName, err := Access.JwtAuth(c)
		if err != nil {
			c.JSON(200, err.Error())
			c.Abort()
		}
		c.Set("username", userName)
		c.Next()
	}
}
