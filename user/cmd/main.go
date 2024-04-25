package main

import (
	"fmt"
	routeV1 "github.com/LucienVen/go-web-demo/user/api/route/v1"
	"github.com/LucienVen/go-web-demo/user/bootstrap"
	logger "github.com/LucienVen/go-web-demo/user/internal/log"
	"github.com/LucienVen/go-web-demo/utils/consul"
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
		lis, err := net.Listen("tcp", env.ServicePort)
		if err != nil {
			errChan <- err
		}

		err = rpc.Serve(lis)
		if err != nil {
			errChan <- err
		}
	}()

	//addr, err := common.GetOutboundIP()
	//if err != nil {
	//	logger.Error("main GetOutboundIP err", zap.Any("err", err.Error()))
	//}

	discoveryConfig := consul.DiscoveryConfig{
		Name:    env.AppName,
		Tags:    env.ConsulSetting.Tag,
		Port:    env.ServicePort,
		Address: env.ConsulSetting.Address,
	}

	discoveryConfig.BuildId()

	consulClient, err := consul.NewConsul(discoveryConfig)
	if err != nil {
		panic(fmt.Sprintf("Run consul.NewConsul err: %v", err.Error()))
	}

	consulClient.RegisterService()


	select {
	case err := <-errChan:
		log.Fatalf("Run Server err: %v", err)
	}
}
