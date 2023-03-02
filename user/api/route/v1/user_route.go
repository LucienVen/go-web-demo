package v1

import (
	"github.com/LucienVen/go-web-demo/user/api/controller"
	"github.com/LucienVen/go-web-demo/user/bootstrap"
	"github.com/LucienVen/go-web-demo/user/domain"
	"github.com/LucienVen/go-web-demo/user/repository"
	"github.com/LucienVen/go-web-demo/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, group *gin.RouterGroup, rpcClient bootstrap.RpcClient) {
	tr := repository.NewUserRepository(db, domain.UserTable)
	tc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(tr, timeout),
		Env:         env,
	}

	group.POST("/user", tc.Create)
	group.GET("/user/:id", tc.GetById)
}
