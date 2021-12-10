package main

import "time"

/*
option 编程思想
把 option 参数封装成一个函数传给我们的目标函数，所有相关的工作由函数来做
非常方便扩展，很多源码都能看到
*/

type Server struct {
	Addr string
	timeout time.Duration
	tls string
}
type config interface {
}
func loadConfig(c interface{}) string {
	return ""
}
func NewServer(addr string, options ...func(*Server)) (*Server, error) {
	srv := &Server{
		Addr:   addr,
	}

	for _, option := range options {
		option(srv)
	}

	return srv, nil
}

func timeout(d time.Duration) func(*Server) {
	return func(srv *Server) {
		srv.timeout = d
	}
}

func tls(c *config) func(*Server) {
	return func(srv *Server) {
		Tls := loadConfig(c)
		srv.tls = Tls
	}
}
