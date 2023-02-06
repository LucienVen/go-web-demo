package domain

import "golang.org/x/net/context"

const (
	TestTable = "user" // 测试用，实质上是user表
)

// 用户状态
type TestStructStatus int64

const (
	TestStructStatus_Unnormal TestStructStatus = iota // 禁用
	TestStructStatus_Normal                           // 正常
)

// ************* Model Struct **************
type TestStruct struct {
	Id         int64            `json:"id" db:"id"`
	Name       string           `json:"name" db:"name"`
	Age        int64            `json:"age" db:"age"`
	Status     TestStructStatus `json:"status"  db:"status"`
	CreateTime int64            `json:"create_time" db:"create_time"`
	UpdateTime int64            `json:"update_time" db:"update_time"`
}

// ************* param **************

// 创建入参
type CreateTestParam struct {
	Name string `json:"name" form:"name" binding:"required"`
	Age  int64  `json:"age" form:"age" binding:"required"`
}

// 修改入参
type UpdateTestParam struct {
	Id         int64            `json:"id"`
	Name       string           `json:"name" form:"name" binding:"required"`
	Age        int64            `json:"age" form:"age" binding:"required"`
	Status     TestStructStatus `json:"status" form:"status"`
	CreateTime int64            `json:"create_time"`
	UpdateTime int64            `json:"update_time"`
}

// 列表查询条件
type TestListParam struct {
	Id       string `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"pageSize"`
}

// 列表查询返回
//type TestListResponse struct {
//	Data []TestStruct `json:"data"`
//}

type TestRepository interface {
	Create(c context.Context, test *TestStruct) (int64, error) // 创建
	List(c context.Context) ([]TestStruct, error)              // 查询列表
	Update(c context.Context, test *TestStruct) error          // 更新
	GetById(c context.Context, id int64) (TestStruct, error)   // 通过ID查询
}

type TestUsecase interface {
	Create(c context.Context, test *TestStruct) error        // 创建
	List(c context.Context) ([]TestStruct, error)            // 查询列表
	Update(c context.Context, test *TestStruct) error        // 更新
	GetById(c context.Context, id int64) (TestStruct, error) // 通过ID查询
}
