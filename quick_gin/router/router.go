package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoshoucheng/hodgepodge/quick_gin/controller/formwork"
	"github.com/zhaoshoucheng/hodgepodge/quick_gin/controller/jump"
	"github.com/zhaoshoucheng/hodgepodge/quick_gin/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Any("/ping",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Use(middleware.Trace())

	v1 := router.Group("/v1")
	formwork.Register(v1)
	jump.Register(v1)

	return router
}
