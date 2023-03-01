package controller

import (
	"github.com/LucienVen/go-web-demo/server2/bootstrap"
	"github.com/LucienVen/go-web-demo/server2/domain"
	logger "github.com/LucienVen/go-web-demo/server2/internal/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
	Env            *bootstrap.Env
}

func (pc *ProductController) Create(c *gin.Context) {
	var err error
	var param domain.ProductParam

	err = c.ShouldBind(&param)
	if err != nil {
		logger.Error("product create:param err", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if ok, err := param.CheckParam(); !ok {
		logger.Error("product create checkParam err", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// TODO 检查 product name 重复

	price := pc.ProductUsecase.GetStorePrice(c, param.Price, domain.ProductPriceMultiple)
	timeNow := time.Now().Unix()

	originData := domain.Product{
		Name:       param.Name,
		Price:      price,
		CreateTime: timeNow,
		UpdateTime: timeNow,
		Status:     domain.ProductStatus_Reviewing, // 新增默认待审核
		IsDelete:   domain.IsDeleteTypes_Normal,
	}

	_, err = pc.ProductUsecase.Create(c, &originData)
	if err != nil {
		logger.Error("product create err", zap.Error(err))
		c.JSON(http.StatusServiceUnavailable, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "",
		Data:    nil,
	})
}
