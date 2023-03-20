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

type Cluster struct {
	HealthChecks []struct {
		Timeout            int `json:"timeout"`
		Interval           int `json:"interval"`
		UnhealthyThreshold int `json:"unhealthy_threshold"`
		HealthyThreshold   int `json:"healthy_threshold"`
		HttpHealthCheck    struct {
			Method string `json:"method"`
			Proto  string `json:"proto"`
			Host   string `json:"host"`
			Path   string `json:"path"`
		} `json:"http_health_check"`
	} `json:"health_checks"`
}

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

	//err = LeaseTest("open", 5)
	//err = WatchTest("open")
	LockTest()
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func getColoring() map[string]interface{} {
	rule1 := make(map[string]interface{})
	rule1["actions"] = map[string]string{"action": "set_group", "value": "gray"}
	rule1["key"] = ""
	rule1["op"] = "equal"
	rule1["type"] = "ip"
	rule1["value"] = "10.99.4.169"
	rule2 := make(map[string]interface{})
	rule2["actions"] = map[string]string{"action": "set_tag", "key": "ruversion", "value": "ruversion_server_test"}
	rule2["key"] = "x-canary-test"
	rule2["op"] = "equal"
	rule2["type"] = "headers"
	rule2["value"] = "test"

	color1 := make(map[string]interface{})
	color1["available_domain"] = []string{"server_test.com", "server_test1.com"}
	color1["name"] = "server_test"
	color1["rules"] = []interface{}{rule1, rule2}
	return color1
}
func getProxyPolicyMatchTags() map[string]interface{} {
	proxy1 := make(map[string]interface{})
	proxy1["apply_on"] = []string{"server_test.com", "server_test1.com"}
	proxy1["enabled_when"] = map[string]interface{}{
		"match_group": "",
		"match_tags":  map[string]string{"ruversion": "ruversion_server_test"},
	}
	proxy1["endpoint_metadata_match"] = map[string]interface{}{
		"app_version": "server_test-v1.0.0",
	}
	proxy1["name"] = "server_test_match_tags"
	return proxy1
}
func getProxyPolicyMatchGroup() map[string]interface{} {
	proxy2 := make(map[string]interface{})
	proxy2["apply_on"] = []string{"server_test.com", "server_test1.com"}
	proxy2["enabled_when"] = map[string]interface{}{
		"match_group": "gray",
		"match_tags":  map[string]string{},
	}
	proxy2["endpoint_metadata_match"] = map[string]interface{}{
		"traffic_strategy": "gray",
	}
	proxy2["name"] = "server_test_match_group"
	return proxy2
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
		fmt.Println(value.ModRevision)
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
