package utils

import (
	"fmt"
	"gin/global"
	"gin/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func GormMySql() *gorm.DB {
	m := global.GCONFIG.MySql
	if m.Name == ""{
		return nil
	}
	var dsn = m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Name + "?parseTime=true&charset=utf8&parseTime=true&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	db,err := gorm.Open(mysql.New(mysqlConfig),&gorm.Config{})
	if err != nil {
		global.HS_LOG.Error("数据库连接失败",zap.Any("err",err))
		return nil
	}
	registerModels(db)

	sql,_:= db.DB()
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	return db
}

func registerModels(db *gorm.DB) {
	err := db.AutoMigrate(
		models.Login{},
		models.Logs{},
		models.AreaCode{},
	)
	if err != nil {
		fmt.Println("表创建失败")
		return
	}
}
