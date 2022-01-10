package proxy

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

//http 代理
func HTTPReverseProxy(c *gin.Context, host string) {
	director := func(req *http.Request) {
		target, err := url.Parse(host)
		if err != nil {
			panic(err)
		}
		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		req.Host = target.Host
		if targetQuery == "" || req.URL.RawQuery == "" {
			buffer := bytes.NewBufferString(targetQuery)
			buffer.WriteString(req.URL.RawQuery)
			req.URL.RawQuery = buffer.String()
		} else {
			buffer := bytes.NewBufferString(targetQuery)
			buffer.WriteString("&")
			buffer.WriteString(req.URL.RawQuery)
			req.URL.RawQuery = buffer.String()
		}
	}

	errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		type Response struct {
			ErrorCode int `json:"errno"`
			ErrorMsg  string       `json:"errmsg"`
			Data      interface{}  `json:"data"`
		}
		resp := &Response{ErrorCode: 99, ErrorMsg: err.Error(), Data: ""}
		c.JSON(200, resp)
		response, _ := json.Marshal(resp)
		c.Set("response", string(response))
		c.AbortWithError(200, err)
	}

	trans := GetTrans(0, 3,10)
	proxy := &httputil.ReverseProxy{Director: director, Transport: trans, ErrorHandler: errFunc}
	proxy.ServeHTTP(c.Writer, c.Request)
	return
}

//组装path
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		buffer := bytes.NewBufferString(a)
		buffer.WriteString(b[1:])
		return buffer.String()
	case !aslash && !bslash:
		buffer := bytes.NewBufferString(a)
		buffer.WriteString("/")
		buffer.WriteString(b)
		return buffer.String()
	}
	buffer := bytes.NewBufferString(a)
	buffer.WriteString(b)
	return buffer.String()
}

func GetTrans(idleNum, connectTimeout, headerTimeout int64) *http.Transport {
	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(connectTimeout) * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext, //3次握手超时设置
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          int(idleNum),
		WriteBufferSize:       1 << 18, //256m
		ReadBufferSize:        1 << 18, //256m
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: time.Duration(headerTimeout) * time.Second, //请求响应超时
	}
	return trans
}
