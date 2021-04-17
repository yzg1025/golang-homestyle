package v1

import (
	"fmt"
	"gin/models"
	"gin/service"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)


func SaveLogs(c *gin.Context)  {
	date := fmt.Sprintf("%v-%v-%v",time.Now().Format("2006"),time.Now().Format("01"),time.Now().Format("02"))
	file,err := ioutil.ReadFile(fmt.Sprintf("./logs/%v.log",date))
	if err != nil {
		fmt.Println(err)
	}
	var Lg models.Logs
	content := string(file)
	_ = c.ShouldBindJSON(&Lg)
	Lg.Content = content
	err = service.SaveLog(Lg)
	if err != nil {
		utils.FailMag("日志存储失败",c)
	}
	utils.SuccessMsg("保存成功",c)
	fmt.Println("file",string(file))

}

