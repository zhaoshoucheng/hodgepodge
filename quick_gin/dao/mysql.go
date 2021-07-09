package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/zhaoshoucheng/hodgepodge/jaeger"
	"time"
)

func Mysql(ctx *gin.Context) interface{}  {
	parentSpan, ok := jaeger.GetSpanFromContext(ctx)
	var span opentracing.Span
	if ok {
		span = opentracing.StartSpan("Mysql",opentracing.ChildOf(parentSpan.Context()))
	} else {
		span = opentracing.StartSpan("Mysql")
	}
	defer span.Finish()
	time.Sleep(time.Second * 3)
	return ""
}

func Redis(ctx *gin.Context) interface{} {
	parentSpan, ok := jaeger.GetSpanFromContext(ctx)
	var span opentracing.Span
	if ok {
		span = opentracing.StartSpan("Redis",opentracing.ChildOf(parentSpan.Context()))
	} else {
		span = opentracing.StartSpan("Redis")
	}
	defer span.Finish()
	time.Sleep(time.Second * 2)
	return ""
}

