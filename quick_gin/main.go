package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhaoshoucheng/hodgepodge/jaeger"
	"github.com/zhaoshoucheng/hodgepodge/quick_gin/conf"
	"github.com/zhaoshoucheng/hodgepodge/quick_gin/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main()  {
	gin.SetMode(gin.TestMode)
	err := jaeger.InitTracer("quick-gin")
	if err != nil {
		panic("jaeger trace init err")
	}
	defer jaeger.Closer()
	port := 8800
	for i := 0; i <= 20; i++ {
		router := router.InitRouter()
		portSte := strconv.Itoa(port + i)
		srv := http.Server{
			Addr:           conf.Host +":"+portSte,
			Handler:        router,
			ReadTimeout:    time.Duration(10) * time.Second,
			WriteTimeout:   time.Duration(10) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()
	}

	quit := make(chan os.Signal)

	// 只监听kill, quit，和ctrl +c
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = ctx
	/*
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	 */
	log.Println("Server exiting")

	signal.Stop(quit)
}
