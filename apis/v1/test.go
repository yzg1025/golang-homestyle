package v1

import (
	"gin/global"
	"gin/models"
	"gin/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
 * we are
 */
func Test(c *gin.Context) {

	d := models.Dog{ID: 1}
	var g []models.GirlGod

	global.HS_DB.Model(&d).Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money > ?", 300)
	}).Association("GirlGods").Find(&g)
	// global.HS_DB.Create(&d)

	utils.SuccessData(g, c)
}
