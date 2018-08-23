package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"ucgo/app/route"
	"ucgo/configs"
	"ucgo/library/iutil"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) //全局设置环境，开发环境为gin.DebugMode，线上环境为gin.ReleaseMode
	// 执行脚本
	if route.Script() {
		return
	}

	// 对外接口
	outerRouter := gin.Default()
	outerRouter.Use(iutil.ParamsMiddleWare())
	route.OuterRouter(outerRouter)
	outerSrv := &http.Server{
		Addr:           ":" + configs.GetConf().Outer_port,
		Handler:        outerRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := outerSrv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
			os.Exit(3)
		}
	}()

	// 对内接口
	innerRouter := gin.Default()
	route.InnerRouter(innerRouter)
	innerSrv := &http.Server{
		Addr:           ":" + configs.GetConf().Inner_port,
		Handler:        innerRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := innerSrv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
			os.Exit(3)
		}
	}()

	// Wait for interrup signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := outerSrv.Shutdown(ctx); err != nil {
		log.Fatal("Outer Server Shutdown:", err)
	}
	if err := innerSrv.Shutdown(ctx); err != nil {
		log.Fatal("Inner Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
