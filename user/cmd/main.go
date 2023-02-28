package main

import (
	"fmt"
	routeV1 "github.com/LucienVen/go-web-demo/server1/api/route/v1"
	"github.com/LucienVen/go-web-demo/server1/bootstrap"
	logger "github.com/LucienVen/go-web-demo/server1/internal/log"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("hello, user")

	app := bootstrap.App()

	env := app.Env
	db := app.Mysql
	rpc := app.Rpc
	rpcClient := app.RpcClient
	defer app.CloseApplication()

	// 初始化日志组件
	logger.InitLogger(env)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	r := gin.Default()

	routerV1 := r.Group("v1")

	routeV1.Setup(env, timeout, db, routerV1, rpcClient)
	errChan := make(chan error)

	// HTTP 服务
	go func() {
		err := r.Run() // 监听并在 0.0.0.0:8080 上启动服务
		if err != nil {
			errChan <- err
		}
	}()

	// 初始化grpc
	go func() {
		lis, err := net.Listen("tcp", env.UserServicePort)
		if err != nil {
			errChan <- err
		}

		err = rpc.Serve(lis)
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		log.Fatalf("Run Server err: %v", err)
	}
}
