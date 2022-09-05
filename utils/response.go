package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

const (
	ERROR   = 199
	SUCCESS = 200
)

func code2String(code int) bool {
	if code == SUCCESS {
		return true
	}
	return false
}

// Result 返回结果
func Result(code bool, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// OK 成功
func OK(c *gin.Context) {
	Result(code2String(SUCCESS), map[string]interface{}{}, "操作成功", c)
}

// SuccessMsg 返回成功信息
func SuccessMsg(msg string, c *gin.Context) {
	Result(code2String(SUCCESS), map[string]interface{}{}, msg, c)
}

// SuccessData 返回成功数据
func SuccessData(data interface{}, c *gin.Context) {
	Result(code2String(SUCCESS), data, "操作成功", c)
}

// SendDetail 返回数据
func SendDetail(data interface{}, msg string, c *gin.Context) {
	Result(code2String(SUCCESS), data, msg, c)
}

// FailMag 失败信息
func FailMag(msg string, c *gin.Context) {
	Result(code2String(ERROR), map[string]interface{}{}, msg, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(code2String(ERROR), data, message, c)
}
