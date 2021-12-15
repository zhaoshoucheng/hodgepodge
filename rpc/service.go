package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)
//http rpc服务
func InitHttpRpc() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//tcp rpc服务
func InitTcpRpc() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

//json rpc服务 同样都是tcp传输，只是用json进行编解码
func InitJsonRpc() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//rpc.ServeConn(conn)
		jsonrpc.ServeConn(conn)
	}
}