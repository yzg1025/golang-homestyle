package main

import (
	"database/sql"
	"fmt"
	"gin/core"
	"gin/global"
)

func main() {
	//cmd.Init()
	db, _ := global.HS_DB.DB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Print("Db Close Fail")
		}
	}(db)
	defer global.HS_REDIS.Close()
	core.RunServer()

	//_, err := os.Create("tt.go")
	//if err != nil {
	//	return
	//}
	//exists, err := utils.PathExists("./generate_code/auto_api")
	//if err != nil {
	//	return
	//}
	//fmt.Println("323", exists)
}
