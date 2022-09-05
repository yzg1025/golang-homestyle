package service

import (
	"gin/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Result struct {
	Weight     string `json:"weight"`
	Categories string `json:"categories"`
	Price      string `json:"price"`
	Desc       string `json:"desc"`
	NickName   string `json:"nick_name"`
	PigId      string `json:"pig_id"`
}

// TestJoins 学习Joins
func TestJoins(c *gin.Context) {
	params := c.Query("id")
	id, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		return
	}
	var total int64
	dog := Result{}
	global.HS_DB.Table("pigs").Debug().
		Select("dogs.categories,dogs.weight,pigs.price,pigs.desc,pigs.pig_id,logins.nick_name").
		Joins("left join cats on pigs.pig_id = cats.cat_id").
		Joins("left join dogs on pigs.pig_id = dogs.dog_id").
		Joins("left join logins on pigs.pig_id = logins.uid").
		Order("pig_id desc").
		Where("pig_id = ?", id).
		Count(&total).
		Scan(&dog)
	c.JSON(http.StatusOK, gin.H{
		"data":  dog,
		"total": total,
		"id":    id,
		"time":  time.Now().Format("2006-01-02 15:01:05"),
	})
}
