package main

import (
	"github.com/zhaoshoucheng/hodgepodge/zlog"
	"time"
)

func main() {
	writer := zlog.MultiLevelWriter(zlog.NewStdoutWriter())
	logger := zlog.New(writer)
	logger.Write([]byte("hello world!"))
	time.Sleep(time.Second * 5)

	return
}

