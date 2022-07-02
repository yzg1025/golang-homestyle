package utils

import (
	"fmt"
	"gin/global"
	"gin/models"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMySql() *gorm.DB {
	m := global.CONFIG.MySql
	if m.Name == "" {
		return nil
	}
	staff := "?parseTime=true&charset=utf8mb4&parseTime=true&loc=Local"
	var dsn = m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Name + staff
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		global.HS_LOG.Error("数据库连接失败", zap.Any("err", err))
		return nil
	}
	registerModels(db)

	sql, _ := db.DB()
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	return db
}

func registerModels(db *gorm.DB) {
	err := db.AutoMigrate(
		models.Login{},
		models.Logs{},
		models.AreaCode{},
		models.Banner{},
		models.PersonalCateGory{},
		models.RoomCount{},
		models.RoomList{},
		models.Location{},
		models.Tag{},

		models.Dog{},
		models.GirlGod{},
		models.Info{},

		models.Albums{},
		models.RankResult{},
		models.RecommendAnchorList{},
		models.RecommendInfoList{},

		models.CategoryAll{},
		models.Categories{},
		models.Subcategories{},

		models.Radio{},

		models.Albumsimple{},
		models.HotWordAlbums{},
		models.TracksAlbum{},

		models.TabList{},
		models.TabViewList{},
		models.HotWordBillboardCategory{},
		models.HotWordBillboard{},

		models.AnchorBasicInfo{},

		models.ScoreDiagram{},
		models.AlbumComments{},
		models.Anchoralbums{},
	)
	if err != nil {
		fmt.Println("表创建失败", err)
		return
	}
}
