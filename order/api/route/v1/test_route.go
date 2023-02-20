package v1

import (
	"github.com/LucienVen/go-web-demo/server2/api/controller"
	"github.com/LucienVen/go-web-demo/server2/bootstrap"
	"github.com/LucienVen/go-web-demo/server2/domain"
	"github.com/LucienVen/go-web-demo/server2/repository"
	"github.com/LucienVen/go-web-demo/server2/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func NewTestRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, group *gin.RouterGroup) {
	tr := repository.NewTestRepository(db, domain.TestTable)
	tc := &controller.TestController{
		TestUsecase: usecase.NewTestUsecase(tr, timeout),
		Env:         env,
	}

	group.POST("/test", tc.Create)
	group.GET("/list/:id", tc.GetById)
}
