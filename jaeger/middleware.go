package jaeger

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"log"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer()
		var span opentracing.Span
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders,opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			span = opentracing.StartSpan(c.Request.URL.Path)
			defer span.Finish()
		} else {
			span = opentracing.StartSpan(c.Request.URL.Path,opentracing.ChildOf(spanCtx))
			defer span.Finish()
		}
		log.Println(c.Request.Header)
		c.Set("ctx-span", span)
		c.Next()
	}
}
