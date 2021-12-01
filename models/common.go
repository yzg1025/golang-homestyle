package models

type Pagination struct {
	PageNo   int `form:"pageNo" json:"pageNo" `    // 页码
	PageSize int `form:"pageSize" json:"pageSize"` // 每页大小
}
