package main

import (
	"database/sql"
	"fmt"
	"gin/core"
	"gin/global"
)

func main() {
	db, _ := global.HS_DB.DB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Print("Db Close Fail")
		}
	}(db)
	defer global.HS_REDIS.Close()
	core.RunServer()
}
