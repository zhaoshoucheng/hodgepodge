package main

import (
	"fmt"
	"github.com/zhaoshoucheng/hodgepodge/zlog"
	"time"
)

func main() {


	timestamp := int64(1641888018)
	t := time.Unix(timestamp, 0)
	fmt.Println(t)
	return
	writer := zlog.MultiLevelWriter(zlog.NewStdoutWriter())
	logger := zlog.New(writer)
	logger.Write([]byte("hello world!"))
	time.Sleep(time.Second * 5)
	return
}

