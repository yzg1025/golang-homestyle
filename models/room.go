package models

type RoomCountParams struct {
	Pagination
	MinLng string `form:"min_lng"`
	MaxLng string `form:"max_lng"`
	MinLat string `form:"min_lat"`
	MaxLat string `form:"max_lat"`
	CLng   string `form:"clng"`
	CLat   string `form:"clat"`
	Zoom   int    `form:"zoom"`
}

type ID struct {
	Name string `from:"name"`
}

type RoomCount struct {
	Code            string  `json:"code"`
	Count           string  `json:"count"`
	SellPrice       int     `gorm:"column:sellPrice";json:"sellPrice"`
	Name            string  `json:"name"`
	Longitude       float32 `json:"longitude"`
	Latitude        float32 `json:"latitude"`
	Touch_lng       string  `json:"touch_lng"`
	Touch_lat       string  `json:"touch_lat"`
	Type            int     `json:"type"`
	Building_status int     `json:"building_status"`
	Distance        int     `json:"distance"`
	Duration        int     `json:"duration"`
	Icon            string  `json:"icon"`
	Zoom            int     `json:"zoom"`
}

type RoomList struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Desc               string  `json:"desc"`
	Apartment_type     int     `json:"apartment_type"`
	Price              int     `json:"price"`
	Price_unit         string  `json:"price_unit"`
	Photo              string  `json:"photo"`
	Photo_alt          string  `json:"photo_alt"`
	Has_video          int     `json:"has_video"`
	Has_3d             int     `json:"has_3d"`
	Commute_info       string  `json:"commute_info"`
	Sale_class         string  `json:"sale_class"`
	Detail_url         string  `json:"detail_url"`
	Sale_status        int     `json:"sale_status"`
	Can_sign_time      int     `json:"can_sign_time"`
	Can_sign_long      int     `json:"can_sign_long"`
	Resblock_id        string  `json:"resblock_id"`
	Resblock_name      string  `json:"resblock_name"`
	Agent_end_date     int     `json:"agent_end_date"`
	Discount_price     int     `json:"discount_price"`
	Sign_date          int     `json:"sign_date"`
	Discounts          int     `json:"discounts"`
	Erebus_sale_status int     `json:"erebus_sale_status"`
	Longitude          float32 `json:"longitude"`
	Latitude           float32 `json:"latitude"`

	Location []Location `gorm:"foreignKey:ID";"column:locations";json:"locations"`
	Tags     []Tag      `gorm:"foreignKey:ID";"column:tags";json:"tags"`
}

type Tag struct {
	ID         string `json:"ld"`
	Title      string `json:"title"`
	Background string `json:"background"`
	Color      string `json:"color"`
}

type Location struct {
	ID   string `json:"id"`
	Name string `json:"title"`
	Url  string `gorm:"column:url";json:"url"`
}

// 测试
type Info struct {
	ID    int
	DogId uint
	Money int
}

type Dog struct {
	ID       int
	Name     string
	Info     Info
	GirlGods []GirlGod `gorm:"many2many:dog_girl_god"`
}

type GirlGod struct {
	ID   int
	Name string
	Dogs []Dog `gorm:"many2many:dog_girl_god"`
}
