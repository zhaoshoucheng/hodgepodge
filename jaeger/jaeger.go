package jaeger

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

const AgentHost = "10.190.33.138:6831"

var TraceCloser io.Closer

func InitTracer(appName string) (err error) {
	cfg :=  config.Configuration{
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: AgentHost,
		},
	}
	TraceCloser, err = cfg.InitGlobalTracer(appName)
	if err != nil {
		return err
	}
	return nil
}

func Closer() {
	TraceCloser.Close()
}

func GetSpanFromContext(ctx *gin.Context) (opentracing.Span, bool) {
	spanFromCtx, exists := ctx.Get("ctx-span")
	if exists {
		span, ok := spanFromCtx.(opentracing.Span)
		if ok {
			return span, true
		}
	}
	return nil, false
}