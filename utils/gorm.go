package utils

import (
	"fmt"
	"gin/global"
	"gin/models"
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
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db,err := gorm.Open(mysql.New(mysqlConfig),&gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
		return nil
	}
	err = db.AutoMigrate(models.Login{})
	if err != nil {
		fmt.Println("表创建失败")
		return nil
	}
	sql,_:= db.DB()
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	return db
}
