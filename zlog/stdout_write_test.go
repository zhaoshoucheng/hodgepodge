package zlog

import (
	"testing"
	"time"
)

func TestNewStdoutWriter(t *testing.T) {

	//在test.go 文件中看不到实时效果是因为go.test会有缓存，在main中就OK
	writer := MultiLevelWriter(NewStdoutWriter())
	logger := New(writer)
	logger.Write([]byte("hello world!"))
	time.Sleep(time.Second * 5)

}