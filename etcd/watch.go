package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//etcd 的watch功能

func WatchTest(env string) (err error) {
	ctx := context.Background()
	cli, err := getEtcdCli(env)
	if err != nil {
		return err
	}
	go func() {
		for {
			cli.Put(ctx, "ping", "pong")
			cli.Delete(ctx, "ping")
			time.Sleep(time.Second)
		}
	}()

	pingVal, err := cli.Get(ctx, "ping")
	if err != nil || len(pingVal.Kvs) == 0 {
		return err
	}
	watchStartRevision := pingVal.Header.Revision + 1
	fmt.Println(watchStartRevision)
	watcher := clientv3.NewWatcher(cli)
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})
	watchRespChan := watcher.Watch(ctx, "ping", clientv3.WithRev(watchStartRevision))
	for watchResp := range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}
	return
}
