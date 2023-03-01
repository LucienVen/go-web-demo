package domain

import (
	"context"
	"github.com/pkg/errors"
)

const (
	ProductTable string = "products" // 产品表
)

const (
	ProductPriceMultiple int64 = 10000
)

type Product struct {
	Id         int64         `json:"id" db:"id"`
	Name       string        `json:"name" db:"name"`   // 产品名称
	Price      int64         `json:"price" db:"price"` // 产品价格（10000倍率）
	CreateTime int64         `json:"create_time" db:"create_time"`
	UpdateTime int64         `json:"update_time" db:"update_time"`
	Status     ProductStatus `json:"status" db:"status"`
	IsDelete   IsDeleteTypes `json:"is_delete" db:"isDelete"` // 软删除
	// TODO tag 分类
}

// ************************************************************

type ProductStatus int64

const (
	ProductStatus_TakeDown  ProductStatus = iota - 1 // 下架
	ProductStatus_Reviewing                          // 待审批
	ProductStatus_Normal                             // 正常
)

// 列表查询
type SearchProductParam struct {
	BaseSearchParam
}

// 产品编辑（新增或修改）
type ProductParam struct {
	Name  string  `json:"name" form:"name"`
	Price float64 `json:"price" form:"price"`
}

func (param *ProductParam) CheckParam() (bool, error) {
	if param.Name == "" {
		return false, errors.New("create param:name cant be nil")
	}

	if param.Price < 0 {
		return false, errors.New("create param:price cant be lt 0")
	}

	return true, nil

}

// 列表返回
type ProductListRes struct {
	Product
}

type ProductRepository interface {
	Create(c context.Context, data *Product) (int64, error)               // 创建
	Update(c context.Context, data *Product) error                        // 修改
	List(c context.Context, param *SearchProductParam) ([]Product, error) // 列表查询
	GetById(c context.Context, id string) (Product, error)                // 列表查询
}

type ProductUsecase interface {
	Create(c context.Context, data *Product) (int64, error)               // 创建
	Update(c context.Context, data *Product) error                        // 修改
	List(c context.Context, param *SearchProductParam) ([]Product, error) // 列表查询
	GetById(c context.Context, id string) (Product, error)                // 列表查询

	GetTruePrice(c context.Context, price int64, Multiplier int64) float64  // 获取真实产品价格（除以倍率）
	GetStorePrice(c context.Context, price float64, Multiplier int64) int64 // 获取存储产品价格（除以倍率）
}
