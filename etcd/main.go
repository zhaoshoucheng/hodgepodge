package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.218.21.167:22379", "10.218.21.196:22379", "10.218.21.204:22379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	//key := "/lkfe/test/filter"
	//key := "/lkfe/test/discover/endpoints/"
	key := "/lkfe/test/"
	resp, err := cli.KV.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithRev(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(resp.Header.Revision)
	//fmt.Println(resp)
	for _, value := range resp.Kvs {
		data := make(map[string]interface{})
		json.Unmarshal(value.Value, &data)
		if _, exists := data["state"]; !exists {
			continue
		}
		fmt.Println(data["state"])
	}

}
