package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/zhaoshoucheng/hodgepodge/gin/conf"
	"github.com/zhaoshoucheng/hodgepodge/gin/docs"
	"github.com/zhaoshoucheng/hodgepodge/gin/router"
	"github.com/zhaoshoucheng/hodgepodge/jaeger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gin.SetMode(gin.DebugMode)
	err := jaeger.InitTracer("quick-gin")
	if err != nil {
		panic("jaeger trace init err")
	}
	defer jaeger.Closer()
	//swagger 配置
	go func() {
		router := router.InitRouter()
		//设置Swagger
		setSwaggerInfo()
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		//服务配置
		srv := http.Server{
			Addr:           conf.Host + ":" + conf.Port1,
			Handler:        router,
			ReadTimeout:    time.Duration(10) * time.Second,
			WriteTimeout:   time.Duration(10) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//主服务
	go func() {
		router := router.InitRouter()
		//服务配置
		srv := http.Server{
			Addr:           conf.Host + ":" + conf.Port2,
			Handler:        router,
			ReadTimeout:    time.Duration(10) * time.Second,
			WriteTimeout:   time.Duration(10) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//一个静态文件服务器
	go func() {
		router := http.FileServer(http.Dir("./"))
		srv := http.Server{
			Addr:           conf.Host + ":" + conf.FileServerPort,
			Handler:        router,
			ReadTimeout:    time.Duration(10) * time.Second,
			WriteTimeout:   time.Duration(10) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

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

func setSwaggerInfo() {
	docs.SwaggerInfo.Title = "my first swagger"
	docs.SwaggerInfo.Description = "my first swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
