package v1

import (
	"fmt"
	"gin/models"
	"gin/service"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func GetPersonalCategory(c *gin.Context) {
   err, list := service.GetAllCatePer()
	fmt.Println("wqw",err,list)
   if err != nil {
   	  utils.FailMag("获取分类失败",c)
	   return
   }
   utils.SuccessData(list,c)
}

func CreatePersonalCategory(c *gin.Context)  {
	var Cate models.PersonalCateGory
	_ = c.ShouldBindJSON(&Cate)
	if Cate.CategoryName == "" {
		utils.FailMag("分类名称不为空",c)
		return
	}
	err := service.CreateCatePer(Cate.CategoryName)
	if err != nil {
		utils.FailMag("创建分类失败",c)
		return
	}
	utils.SuccessMsg("创建分类成功",c)
}