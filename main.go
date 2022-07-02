package main

import (
	"database/sql"
	"fmt"
	v1 "gin/apis/v1"
	"gin/core"
	"gin/global"
	"gin/utils"
)

func main() {
	utils.Init()
	t := v1.NewInterDb("yzg", "32")
	desc := t.MyDesc()
	fmt.Println(desc)

	db, _ := global.HS_DB.DB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	core.RunServer()
}
