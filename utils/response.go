package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int  `json:"code"`
	Data interface{}  `json:"data"`
	Msg  string  `json:"msg"`
}

const (
	ERROR = 1
	SUCCESS = 0
)

// 返回结果
func Result(code int,data interface{},msg string,c *gin.Context)  {
	c.JSON(http.StatusOK,Response{
		code,
		data,
		msg,
	})
}

// 成功
func OK(c *gin.Context)  {
	Result(SUCCESS, map[string]interface{}{},"操作成功",c)
}

// 返回成功信息
func SuccessMsg(msg string,c *gin.Context)  {
	Result(SUCCESS, map[string]interface{}{},msg,c)
}

// 返回成功数据
func SuccessData(data interface{},c *gin.Context)  {
	Result(SUCCESS,data,"c成功",c)
}

// 返回数据
func SendDetail(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS,data,msg,c)
}

// 失败信息
func FailMag(msg string,c *gin.Context)  {
	Result(ERROR, map[string]interface{}{},msg,c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}


