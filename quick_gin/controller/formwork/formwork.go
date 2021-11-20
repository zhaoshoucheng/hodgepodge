package formwork

import "github.com/gin-gonic/gin"

type FormworkController struct {
}

// helloworld godoc
// @Summary helloworld
// @Description 测试服务连通性helloworld
// @Tags V1
// @ID /v1/do1
// @Accept json
// @Produce  json
// @Success 200 {object} string "success"
// @Router /v1/do1 [post]
func (fw *FormworkController) Test(ctx *gin.Context) {
	ctx.JSON(200,"hello world")
}

func Register(eng *gin.RouterGroup) {
	fw := &FormworkController{}
	eng.POST("/do1",fw.Test)
	eng.GET("/do2",fw.Test)
}
func Register2(eng *gin.RouterGroup) {
	fw := &FormworkController{}
	eng.POST("",fw.Test)
	eng.GET("",fw.Test)
}