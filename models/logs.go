package models

type Logs struct {
	ID         float32 `json:"id"`
	Level      string  `json:"level"`
	Time       string  `json:"time"`
	Msg        string  `json:"msg"`
	Err        string  `json:"err"`
	MethodName string  `gorm:"column:methodName" json:"methodName"`
}
