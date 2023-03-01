package domain

const (
	OrderTable string = "order" // order表
)

const (
	OrderPaymentMultiple int64 = 10000 // 金额存储转换倍数
)

type Order struct {
	Id         int64       `json:"id" db:"id"`
	UserId     int64       `json:"user_id" db:"user_id"`       // 用户ID
	ProductId  int64       `json:"product_id" db:"product_id"` // 产品ID
	Payment    int64       `json:"payment" db:"payment"`       // 实际支付（10000倍数）
	Status     OrderStatus `json:"status" db:"status"`         // 订单状态
	CreateTime int64       `json:"create_time" db:"create_time"`
	UpdateTime int64       `json:"update_time" db:"update_time"`
}

// ****************************** ******************************

type OrderStatus int64 // 订单状态
const (
	OrderStatus_Cancel  OrderStatus = iota - 1 // 取消
	OrderStatus_Pending                        // 待确认
	OrderStatus_Success                        // 成功
)
