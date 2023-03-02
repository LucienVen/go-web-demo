package usecase

import (
	"context"
	"github.com/LucienVen/go-web-demo/user/domain"
	"time"
)

type UserUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

// 创建
func (tu *UserUsecase) Create(c context.Context, user *domain.UserStruct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	_, err := tu.userRepository.Create(ctx, user)
	return err
}

// 查询列表
func (tu *UserUsecase) List(c context.Context) ([]domain.UserStruct, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.userRepository.List(ctx)
}

// 更新
func (tu *UserUsecase) Update(c context.Context, user *domain.UserStruct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.userRepository.Update(ctx, user)
}

func (tu UserUsecase) GetById(c context.Context, id int64) (domain.UserStruct, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.userRepository.GetById(ctx, id)
}
