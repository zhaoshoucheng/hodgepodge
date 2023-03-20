package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// 分布式锁的实现

type Lock struct {
	lease      clientv3.Lease
	leaseId    clientv3.LeaseID
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func (l *Lock) Lock() (lock bool, err error) {
	cli, err := getEtcdCli("open")
	if err != nil {
		return false, err
	}
	l.lease = clientv3.NewLease(cli)
	l.ctx, l.cancelFunc = context.WithCancel(context.TODO())
	leaseGrant, err := l.lease.Grant(context.TODO(), 5)
	if err != nil {
		return false, err
	}
	l.leaseId = leaseGrant.ID
	kv := clientv3.NewKV(cli)
	txn := kv.Txn(l.ctx)
	txn.If(clientv3.Compare(clientv3.CreateRevision("lock"), "=", 0)).
		Then(clientv3.OpPut("lock", "g", clientv3.WithLease(l.leaseId)))
	txnResp, err := txn.Commit()
	if err != nil {
		return false, err
	}
	if !txnResp.Succeeded {
		return false, nil
	}
	//自动续约
	keepRespChan, err := l.lease.KeepAlive(l.ctx, l.leaseId)
	_ = keepRespChan
	/*
		go func() {
			for {
				select {
				case keepResp := <-keepRespChan:
					if keepResp == nil {
						fmt.Println("租约已经失效了")
						goto END
					} else { // 每秒会续租一次, 所以就会受到一次应答
						fmt.Println("收到自动续租应答:", keepResp.ID)
					}
				}
			}
		END:
		}()
	*/

	return true, nil
}
func (l *Lock) Unlock() {
	//l.cancelFunc()
	l.lease.Revoke(l.ctx, l.leaseId)
	/*
		cli, _ := getEtcdCli("open")
		value, _ := cli.Get(context.Background(), "lock")
		if len(value.Kvs) == 0 {
			fmt.Println("lock del")
		} else {
			fmt.Println("get lock ", string(value.Kvs[0].Value))
		}
	*/

}

func LockTest() {
	go Node("node1", 5)
	go Node("node2", 3)
	select {}
}

func Node(node string, t time.Duration) {
	l := Lock{}
	for {
		getLock, err := l.Lock()
		if err != nil || !getLock {
			continue
		}
		fmt.Println("i get the lock: ", node)
		time.Sleep(time.Second * t)
		l.Unlock()
		fmt.Println("i release the lock: ", node)
		time.Sleep(time.Second)
	}
}
