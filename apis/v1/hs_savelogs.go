package v1

import (
	"encoding/json"
	"fmt"
	"gin/global"
	"gin/models"
	"gin/service"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"strings"
	"time"
)

type LogManger interface {
	GetLogs(c *gin.Context)
	SaveLogs(c *gin.Context)
}

type LogCenter struct {
	lg service.UserServiceCenter
}

func NewLogs() *LogCenter {
	return &LogCenter{}
}

func (log *LogCenter) GetLogs(c *gin.Context) {
	err, list := log.lg.GetLog()
	if err != nil {
		global.HS_LOG.Error("获取失败", zap.Any("err", err))
		utils.FailMag("获取失败", c)
	}
	utils.SuccessData(list, c)
}

func (log *LogCenter) SaveLogs(c *gin.Context) {
	date := fmt.Sprintf(
		"%v-%v-%v",
		time.Now().Format("2006"),
		time.Now().Format("01"),
		time.Now().Format("02"),
	)
	file, err := ioutil.ReadFile(fmt.Sprintf("./logs/%v.log", date))
	if err != nil {
		fmt.Println(err)
	}

	content := string(file)
	replaceContents := strings.ReplaceAll(content, "}", "}|")
	contentsToArr := strings.Split(replaceContents, "|")

	var logs []models.Logs
	_ = c.ShouldBindJSON(&logs)
	for index, val := range contentsToArr[:len(contentsToArr)-1] {
		var log models.Logs
		err = json.Unmarshal([]byte(val), &log)
		if err != nil {
			fmt.Print("error", err)
		}
		newLog := &models.Logs{
			ID:         float32(index),
			Time:       log.Time,
			Err:        log.Err,
			MethodName: log.MethodName,
			Msg:        log.Msg,
			Level:      log.Level,
		}
		logs = append(logs, *newLog)
	}
	err = log.lg.SaveLog(logs)
	if err != nil {
		utils.FailMag("err", c)
		return
	}
	utils.SuccessMsg("日志存储成功", c)
}
