package jump

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoshoucheng/hodgepodge/quick_gin/conf"
	"github.com/zhaoshoucheng/hodgepodge/quick_gin/dao"
	"github.com/zhaoshoucheng/hodgepodge/quick_gin/util"
	"log"
	"time"
)
const pathA = "/serverA"
const pathAB = "/serverA/to/serverB"
const pathAC = "/serverA/to/serverC"
const pathBD = "/serverB/to/serverD"
const pathD = "/serverD"


type JumpController struct {
}

func (JC *JumpController) A(ctx *gin.Context) {
	//A -> b
	abURl := "http://"+conf.Host+":"+conf.Port1+"/v1"+pathAB
	respAB, err := util.Get(ctx, abURl)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(respAB)
	}
	//A -> c
	abUrl := "http://"+conf.Host+":"+conf.Port1+"/v1"+pathAC
	respAC, err := util.Get(ctx, abUrl)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(respAC)
	}
	ctx.JSON(200,"serverA"+respAB+respAC)
}

func (JC *JumpController) AB(ctx *gin.Context) {
	//b -> d
	bdUrl := "http://"+conf.Host+":"+conf.Port1+"/v1"+pathBD
	resp, err := util.Get(ctx, bdUrl)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
	//do sth
	time.Sleep(time.Second)
	ctx.JSON(200,"serverAB" + resp)
}
func (JC *JumpController) AC(ctx *gin.Context) {
	_ = dao.Mysql(ctx)
	ctx.JSON(200,"serverAC")
}
func (JC *JumpController) BD(ctx *gin.Context) {
	// ->d
	url := "http://"+conf.Host+":"+conf.Port1+"/v1"+pathD
	resp, err := util.Get(ctx, url)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
	ctx.JSON(200,"serverBD"+resp)
}
func (JC *JumpController) D(ctx *gin.Context) {
	_ = dao.Mysql(ctx)
	_ = dao.Redis(ctx)
	ctx.JSON(200,"serverD")
	//span,ok := jaeger.GetSpanFromContext(ctx)
	//if !ok {
	//	return
	//}
	//span1 := opentracing.StartSpan("serverD-1",opentracing.ChildOf(span.Context()))
	//defer span1.Finish()
	//span2 := opentracing.StartSpan("serverD-2",opentracing.FollowsFrom(span.Context()))
	//defer span2.Finish()
}
func Register(eng *gin.RouterGroup) {
	jc := &JumpController{}
	eng.GET(pathA,jc.A)
	eng.GET(pathAB,jc.AB)
	eng.GET(pathAC,jc.AC)
	eng.GET(pathBD,jc.BD)
	eng.GET(pathD,jc.D)
}
