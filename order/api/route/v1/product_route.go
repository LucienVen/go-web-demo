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

func NewProductRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, group *gin.RouterGroup) {
	pr := repository.NewProductRepository(db, domain.ProductTable)
	pc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr, timeout),
		Env:            env,
	}

	group.POST("/product", pc.Create)
}
