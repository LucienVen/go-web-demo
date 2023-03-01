package v1

import (
	"github.com/LucienVen/go-web-demo/order/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

//func Handle(router *gin.Engine) {
//	group := router.Group("/v1")
//	// TODO 暂时
//	testController := new(controller.TestController)
//	group.GET("/hello", testController.Hello)
//	group.GET("/call-my-name", testController.CallMyName)
//
//}

func Setup(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, routeV1 *gin.RouterGroup) {
	// public APIs 公共API
	publicRouteV1 := routeV1.Group("")
	NewTestRouter(env, timeout, db, publicRouteV1)
	NewProductRouter(env, timeout, db, publicRouteV1)

	// Private APIs 私有API
	//protectedRouteV1 := routeV1.Group("")
	// 中间件
}
