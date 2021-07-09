package util

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/zhaoshoucheng/hodgepodge/jaeger"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Get(ctx *gin.Context,url string) (string,error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	span, ok := jaeger.GetSpanFromContext(ctx)
	if ok {
		log.Println(span)
		err = opentracing.GlobalTracer().Inject(span.Context(),opentracing.HTTPHeaders,opentracing.HTTPHeadersCarrier(req.Header))
		if err != nil {
			log.Println(err)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}
