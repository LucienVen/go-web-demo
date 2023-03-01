package usecase

import (
	"context"
	"github.com/LucienVen/go-web-demo/order/domain"
	"time"
)

type ProductUsecase struct {
	productRepository domain.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUsecase(pr domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &ProductUsecase{
		productRepository: pr,
		contextTimeout:    timeout,
	}
}

func (pu *ProductUsecase) Create(c context.Context, data *domain.Product) (int64, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	thisId, err := pu.productRepository.Create(ctx, data)
	return thisId, err
}

func (pu *ProductUsecase) Update(c context.Context, data *domain.Product) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	err := pu.productRepository.Update(ctx, data)
	return err
}

func (pu *ProductUsecase) List(c context.Context, param *domain.SearchProductParam) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	ret, err := pu.productRepository.List(ctx, param)
	return ret, err
}

func (pu *ProductUsecase) GetById(c context.Context, id string) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	ret, err := pu.productRepository.GetById(ctx, id)
	return ret, err
}

func (pu *ProductUsecase) GetTruePrice(c context.Context, price int64, multi int64) float64 {
	return float64(price) / float64(multi)
}

func (pu *ProductUsecase) GetStorePrice(c context.Context, price float64, multi int64) int64 {
	return int64(price * float64(multi))
}
