package v1

import (
	"github.com/LucienVen/go-web-demo/order/api/controller"
	"github.com/LucienVen/go-web-demo/order/bootstrap"
	"github.com/LucienVen/go-web-demo/order/domain"
	"github.com/LucienVen/go-web-demo/order/repository"
	"github.com/LucienVen/go-web-demo/order/usecase"
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
