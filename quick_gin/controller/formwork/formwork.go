package formwork

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhaoshoucheng/hodgepodge/proxy"
)

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
func (fw *FormworkController) TestOne(ctx *gin.Context) {
	fmt.Println(ctx.Request.Host)
	if ctx.Request.Host == "127.0.0.1:8888" {
		//代理转发
		proxy.HTTPReverseProxy(ctx, "http://127.0.0.1:8898")
		return
	}
	ctx.JSON(200,fmt.Sprintf("hello world one %s",ctx.Request.Host))
}

func (fw *FormworkController) TestTwo(ctx *gin.Context) {
	ctx.JSON(200,fmt.Sprintf("hello world two %s",ctx.Request.Host))
}

func (fw *FormworkController) Test(ctx *gin.Context) {
	ctx.JSON(200,fmt.Sprintf("hello world %s",ctx.Request.Host))
}

func Register(eng *gin.RouterGroup) {
	fw := &FormworkController{}
	eng.POST("/do1",fw.TestOne)
	eng.GET("/do2",fw.TestTwo)
}
func Register2(eng *gin.RouterGroup) {
	fw := &FormworkController{}
	eng.POST("",fw.Test)
	eng.GET("",fw.Test)
}