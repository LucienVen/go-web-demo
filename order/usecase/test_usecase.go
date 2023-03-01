package usecase

import (
	"context"
	"github.com/LucienVen/go-web-demo/order/domain"
	"time"
)

type TestUsecase struct {
	testRepository domain.TestRepository
	contextTimeout time.Duration
}

func NewTestUsecase(testRepository domain.TestRepository, timeout time.Duration) domain.TestUsecase {
	return &TestUsecase{
		testRepository: testRepository,
		contextTimeout: timeout,
	}
}

// 创建
func (tu *TestUsecase) Create(c context.Context, test *domain.TestStruct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	_, err := tu.testRepository.Create(ctx, test)
	return err
}

// 查询列表
func (tu *TestUsecase) List(c context.Context) ([]domain.TestStruct, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.testRepository.List(ctx)
}

// 更新
func (tu *TestUsecase) Update(c context.Context, test *domain.TestStruct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.testRepository.Update(ctx, test)
}

func (tu TestUsecase) GetById(c context.Context, id int64) (domain.TestStruct, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	return tu.testRepository.GetById(ctx, id)
}

//List(c context.Context) ([]TestStruct, error)
//Update(c context.Context, test *TestStruct) error        // 更新
//GetById(c context.Context, id int64) (TestStruct, error) // 通过ID查询
