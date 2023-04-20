package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"os"
	"strings"
	"time"
)

var MyHostName string

type etcdLeader struct {
	sess *concurrency.Session
	elec *concurrency.Election
}

func process(processID string) {
	cli, err := getEtcdCli("open")
	if err != nil {
		fmt.Println(err)
		panic(processID)
	}
	ctx := context.Background()
	maintainer, err := Campaign(context.Background(), cli, "process", "")
	if err != nil {
		fmt.Println("process " + processID + " Campaign err" + err.Error())
		return
	}
	fmt.Println(processID + "is maintainer")
	ticker := time.NewTicker(time.Second)
	count := 0
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			err = maintainer.Resign(context.Background())
			if err != nil {
				fmt.Printf("error occured when resign lfe_sync: %v", err)
			}
			return
		case <-ticker.C:
			err = maintainer.Proclaim(ctx, "")
			if err != nil {
				fmt.Printf("error occured when proclaim lfe_sync: %v", err)
				maintainer.Resign(ctx)
				return
			}
			count++
			//do sth
			fmt.Println("processID ", processID+" ", count)
			if count == 5 {
				err = maintainer.Resign(context.Background())
				if err != nil {
					fmt.Printf("error occured when resign lfe_sync: %v", err)
				}
				return
			}
		}
	}
}
func MasterTest() error {
	go process("1")
	go process("2")
	select {}

}

func Campaign(ctx context.Context, client *clientv3.Client, prefix, value string) (*etcdLeader, error) {
	if value == "" {
		value = fmt.Sprintf("%s-%d", MyHostName, timestampMs())
	}
	s, err := concurrency.NewSession(client)
	if err != nil {
		return nil, fmt.Errorf("failed to generate session: %v", err)
	}
	prefix = "/openresty/" + "-concurrency/" + strings.TrimPrefix(prefix, "/")

	elec := concurrency.NewElection(s, prefix)
	return &etcdLeader{
		sess: s,
		elec: elec,
	}, elec.Campaign(ctx, value) // blocked until elected
}

func (l *etcdLeader) Proclaim(ctx context.Context, value string) error {
	if l.elec == nil {
		return fmt.Errorf("already closed")
	}
	if value == "" {
		value = fmt.Sprintf("%s-%d", MyHostName, timestampMs())
	}
	return l.elec.Proclaim(ctx, value)
}

func (l *etcdLeader) Resign(ctx context.Context) error {
	if l.elec == nil {
		return nil
	}
	defer l.sess.Close()
	defer func() {
		l.elec = nil
		l.sess = nil
	}()
	return l.elec.Resign(ctx)
}

func init() {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalf("failed to get hostname: %v", err)
	}
	MyHostName = hn
}
func timestampMs() int64 {
	return time.Now().UnixNano() / 1e6
}
