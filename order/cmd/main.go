package main

import (
	"fmt"
	routeV1 "github.com/LucienVen/go-web-demo/server2/api/route/v1"
	"github.com/LucienVen/go-web-demo/server2/bootstrap"
	"github.com/LucienVen/go-web-demo/server2/internal/log"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	fmt.Println("hello, order")

	app := bootstrap.App()

	env := app.Env
	db := app.Mysql
	defer app.CloseApplication()

	// 初始化日志组件
	log.InitLogger(env)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	r := gin.Default()

	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "hello, order",
	//	})
	//})

	//routeV1.Handle(r)

	routerV1 := r.Group("v1")

	routeV1.Setup(env, timeout, db, routerV1)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
