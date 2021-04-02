package main

import (
	"fmt"
	"gin/global"
	"gin/utils"
)
func main() {
	utils.Init()
	fmt.Println("2323",global.GCONFIG.MySql)
}
