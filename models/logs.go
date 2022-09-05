package models

type Logs struct {
	ID         float32 `json:"id"`
	Level      string  `json:"level"`
	Time       string  `json:"time"`
	Msg        string  `json:"msg"`
	Err        string  `json:"err"`
	MethodName string  `gorm:"column:methodName" json:"methodName"`
}

type Cat struct {
	Color string `json:"color"`
	Name  string `json:"name"`
	CatId int64  `json:"catId"`
}

type Dog struct {
	Sex        string `json:"sex"`
	DogId      int64  `json:"dogId"`
	Weight     int64  `json:"weight"`
	Categories string `json:"categories"`
}

type Pig struct {
	PigId int64   `json:"pig_id"`
	Price float64 `json:"price"`
	Color string  `json:"color"`
	Desc  string  `json:"desc"`
}

type Duck struct {
	DuckId int64 `json:"duck_id"`
}
