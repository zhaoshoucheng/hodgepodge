package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//etcd 租约与续约实践

func LeaseTest(env string, ttl int64) (err error) {
	cli, err := getEtcdCli(env)
	if err != nil {
		return
	}
	lease := clientv3.NewLease(cli)
	leaseGrant, err := lease.Grant(context.Background(), ttl)
	if err != nil {
		return
	}

	if _, err = cli.Put(context.Background(), "ping", "pong", clientv3.WithLease(leaseGrant.ID)); err != nil {
		return
	}
	/*
		保持长链接，每s续租一次
	*/
	keepRespChan, err := lease.KeepAlive(context.TODO(), leaseGrant.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		//查看续期情况
		for {
			select {
			case resp := <-keepRespChan:
				if resp == nil {
					fmt.Println("租约失效")
					return
				} else {
					fmt.Println("租约成功", resp)
				}
			}
		}
	}()

	for {
		values, err := cli.Get(context.Background(), "ping")
		if err != nil {
			break
		}
		if values.Count == 0 {
			fmt.Println("已经过期")
		}
		fmt.Println("没过期", values.Kvs)
		time.Sleep(time.Second * 1)
	}
	time.Sleep(5 * time.Second)
	return
}
