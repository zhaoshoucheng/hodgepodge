package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	"log"
	"net"
	"time"
)

func main() {
	configFile := flag.String("c", "config", "config filename (extension in include)")
	flag.Parse()
	viper.SetConfigName(*configFile)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
		return
	}

	PrintlnAllValue("/openresty/", "itest")
	return
}

func PrintlnAllKey(path string, env string) error {
	resp, err := getValue(path, env)
	if err != nil {
		return err
	}
	for _, value := range resp.Kvs {
		fmt.Println(string(value.Key))
	}
	return nil
}

func PrintlnAllValue(path string, env string) (values []string, err error) {
	resp, err := getValue(path, env)
	if err != nil {
		return
	}
	for _, value := range resp.Kvs {
		fmt.Println(string(value.Value))
		values = append(values, string(value.Value))
	}
	return
}

func getValue(path string, env string) (*clientv3.GetResponse, error) {
	cli, err := getEtcdCli(env)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("connect to etcd failed, err:%v\n", err))
	}
	ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	resp, err := cli.KV.Get(ctx, path, clientv3.WithPrefix(), clientv3.WithRev(0))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DeleteValue(path string, env string) error {
	cli, err := getEtcdCli(env)
	if err != nil {
		return errors.New(fmt.Sprintf("connect to etcd failed, err:%v\n", err))
	}
	ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	resp, err := cli.KV.Delete(ctx, path)
	if err != nil {
		return err
	}
	fmt.Println(resp.Deleted)
	return nil
}

func PutValue(path string, env string, value string) error {
	cli, err := getEtcdCli(env)
	if err != nil {
		return errors.New(fmt.Sprintf("connect to etcd failed, err:%v\n", err))
	}
	ctx, _ := context.WithTimeout(context.TODO(), time.Second)
	resp, err := cli.KV.Put(ctx, path, value)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}
func getEtcdCli(env string) (*clientv3.Client, error) {
	endpoints := viper.GetStringSlice(env + ".endpoints")
	if len(endpoints) == 0 {
		return nil, errors.New("config err")
	}
	userName := viper.GetString(env + ".username")
	passWord := viper.GetString(env + ".password")
	return clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 10 * time.Second,
		Username:    userName,
		Password:    passWord,
	})
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err.Error()
	}
	var str string
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				str = str + ipnet.IP.String()
			}
		}
	}
	return str
}
