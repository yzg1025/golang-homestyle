package v1

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context)  {
	utils.SuccessMsg("测试成功",c)
}
