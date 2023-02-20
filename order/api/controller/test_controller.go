package controller

import "C"
import (
	"github.com/LucienVen/go-web-demo/server2/bootstrap"
	"github.com/LucienVen/go-web-demo/server2/domain"
	logger "github.com/LucienVen/go-web-demo/server2/internal/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type TestController struct {
	TestUsecase domain.TestUsecase
	Env         *bootstrap.Env
}

func (tc *TestController) Create(c *gin.Context) {
	var testParam domain.CreateTestParam

	err := c.ShouldBind(&testParam)
	if err != nil {
		logger.Error("testController create err", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	timeNow := time.Now().Unix()

	testData := domain.TestStruct{
		Name:       testParam.Name,
		Age:        testParam.Age,
		Status:     domain.TestStructStatus_Normal,
		CreateTime: timeNow,
		UpdateTime: timeNow,
	}

	err = tc.TestUsecase.Create(c, &testData)
	if err != nil {
		logger.Error("[testController] create err", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    nil,
	})
}

func (tc *TestController) GetById(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		logger.Error("[testController] GetById param is null")
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "id should not be null"})
		return
	}

	id, _ := strconv.ParseInt(idParam, 10, 64)
	testData, err := tc.TestUsecase.GetById(c, id)
	if err != nil {
		logger.Error("[testController] usecase GetById err", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    testData,
	})
}
