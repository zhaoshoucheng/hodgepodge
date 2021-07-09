package formwork

import "github.com/gin-gonic/gin"

type FormworkController struct {
}

func (fw *FormworkController) Test(ctx *gin.Context) {
	ctx.JSON(200,"hello world")
}

func Register(eng *gin.RouterGroup) {
	fw := &FormworkController{}
	eng.POST("/do1",fw.Test)
	eng.GET("/do2",fw.Test)
}