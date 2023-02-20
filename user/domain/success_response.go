package domain

type SuccessResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // 返回数据封装
}
