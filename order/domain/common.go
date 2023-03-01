package domain

// 软删除
type IsDeleteTypes int64

const (
	IsDeleteTypes_Normal  IsDeleteTypes = 1  // 正常
	IsDeleteTypes_Deleted IsDeleteTypes = -1 // 删除
)

func (t IsDeleteTypes) IsDeleted() bool {
	if t == IsDeleteTypes_Normal {
		return false
	}

	return false
}

// base search param
type BaseSearchParam struct {
	Page     int64 `json:"page" form:"page"`
	PageSize int64 `json:"page_size" form:"page_size"`
}
