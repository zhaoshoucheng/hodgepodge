package main

import (
	"fmt"
	"github.com/zhaoshoucheng/hodgepodge/zlog"
	"time"
)

func main() {

	var lng int64
	lng = 1234567
	lngStr := fmt.Sprintf("%f",float64(lng)/100000)
	fmt.Println(lngStr)
	return
	type test struct {
		ID string
	}
	test1 := &test{
		ID: "1",
	}
	test2 := &test{
		ID: "2",
	}
	test3 := &test{
		ID: "3",
	}
	testList := []*test{test1,test2,test3}
	testMap := make(map[string]*test)
	for _, testItem := range testList {
		testMap[testItem.ID] = testItem
	}
	for key ,value := range testMap {
		fmt.Println(key, "------", value)
	}


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

