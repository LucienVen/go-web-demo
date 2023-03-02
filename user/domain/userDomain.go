package domain

import (
	"golang.org/x/net/context"
)

const (
	UserTable = "user" // 测试用，实质上是user表
)

// 用户状态
type UserStructStatus int64

const (
	UserStructStatus_Unnormal UserStructStatus = iota // 禁用
	UserStructStatus_Normal                           // 正常
)

// ************* Model Struct **************
// 直接关联数据库（可以看做与数据库字段的映射）
type UserStruct struct {
	Id         int64            `json:"id" db:"id"`
	Name       string           `json:"name" db:"name"`
	Age        int64            `json:"age" db:"age"`
	Status     UserStructStatus `json:"status"  db:"status"`
	CreateTime int64            `json:"create_time" db:"create_time"`
	UpdateTime int64            `json:"update_time" db:"update_time"`
}

// ************* param **************
// 入参出参定义与model分开

// 创建入参
type CreateUserParam struct {
	Name string `json:"name" form:"name" binding:"required"`
	Age  int64  `json:"age" form:"age" binding:"required"`
}

// 修改入参
type UpdateUserParam struct {
	Id         int64            `json:"id"`
	Name       string           `json:"name" form:"name" binding:"required"`
	Age        int64            `json:"age" form:"age" binding:"required"`
	Status     UserStructStatus `json:"status" form:"status"`
	CreateTime int64            `json:"create_time"`
	UpdateTime int64            `json:"update_time"`
}

// 列表查询条件
type UserListParam struct {
	Id       string `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"pageSize"`
}

// 列表查询返回
//type UserListResponse struct {
//	Data []UserStruct `json:"data"`
//}

type UserRepository interface {
	Create(c context.Context, user *UserStruct) (int64, error) // 创建
	List(c context.Context) ([]UserStruct, error)              // 查询列表
	Update(c context.Context, user *UserStruct) error          // 更新
	GetById(c context.Context, id int64) (UserStruct, error)   // 通过ID查询
}

type UserUsecase interface {
	Create(c context.Context, user *UserStruct) error        // 创建
	List(c context.Context) ([]UserStruct, error)            // 查询列表
	Update(c context.Context, user *UserStruct) error        // 更新
	GetById(c context.Context, id int64) (UserStruct, error) // 通过ID查询
}

//type UserRpcUsecase interface {
//	SayHello(c context.Context, req *pb.SayHelloReq) (*pb.SayHelloRes, error)
//}
