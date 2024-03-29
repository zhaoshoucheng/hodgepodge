package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoshoucheng/hodgepodge/Access"
	"github.com/zhaoshoucheng/hodgepodge/gin/middleware"
)

func token(c *gin.Context) {
	token, err := Access.Tokens(c)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, token)
}

func getUserName(c *gin.Context) {
	if username, exists := c.Get("username"); exists {
		c.JSON(200, username)
		return
	}
	c.JSON(200, "username not exists")
}
func Register(eng *gin.RouterGroup) {
	eng.GET("/oauth/tokens", token)

	eng.Use(middleware.JwtAuthMiddleware())
	eng.GET("/username", getUserName)
}
