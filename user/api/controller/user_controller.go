package controller

import "C"
import (
	"github.com/LucienVen/go-web-demo/user/bootstrap"
	"github.com/LucienVen/go-web-demo/user/domain"
	logger "github.com/LucienVen/go-web-demo/user/internal/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

func (tc *UserController) Create(c *gin.Context) {
	var userParam domain.CreateUserParam

	err := c.ShouldBind(&userParam)
	if err != nil {
		logger.Error("userController create err", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	timeNow := time.Now().Unix()

	userData := domain.UserStruct{
		Name:       userParam.Name,
		Age:        userParam.Age,
		Status:     domain.UserStructStatus_Normal,
		CreateTime: timeNow,
		UpdateTime: timeNow,
	}

	err = tc.UserUsecase.Create(c, &userData)
	if err != nil {
		logger.Error("[userController] create err", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    nil,
	})
}

func (tc *UserController) GetById(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		logger.Error("[userController] GetById param is null")
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "id should not be null"})
		return
	}

	id, _ := strconv.ParseInt(idParam, 10, 64)
	userData, err := tc.UserUsecase.GetById(c, id)
	if err != nil {
		logger.Error("[userController] usecase GetById err", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    userData,
	})
}
