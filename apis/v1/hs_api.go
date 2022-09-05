package v1

import (
	"fmt"
	"gin/generate_code/auto_model"
	"gin/global"
	"gin/models"
	"gin/service"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type ApiSer struct {
	serv service.ApiService
	api  models.SimpleAPI
}

type ApiInter interface {
	CreateApi(c *gin.Context)
	QueryApi(c *gin.Context)
}

func NewAPIInter() *ApiSer {
	return &ApiSer{}
}

func (a *ApiSer) CreateApi(c *gin.Context) {
	_ = c.ShouldBindJSON(&a.api)
	if err := utils.Verify(a.api, utils.APIVar); err != nil {
		global.HS_LOG.Error("api 参数不为空", zap.Any("err", err), zap.Any("method", "CreateApi"))
		utils.FailMag(err.Error(), c)
		return
	}
	// 生成model
	paths := a.TemplatePath()
	for _, path := range paths {
		a.AutoCode(path)
	}

	a.api.CreateTime = time.Now()
	if a.api.AccessAuth {
		a.api.ApiPath = "/api/hs_privacy" + strings.ToLower(a.api.ApiMethod)
	} else {
		a.api.ApiPath = "/api/hs_base" + strings.ToLower(a.api.ApiMethod)
	}
	errs := service.CreateApi(a.api)
	if errs != nil {
		utils.FailMag("添加失败", c)
		return
	}

	utils.SuccessMsg("添加成功", c)
}

func (a *ApiSer) QueryApi(c *gin.Context) {
	page, _ := strconv.Atoi(c.PostForm("page"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	apiName := c.PostForm("apiName")
	apiPath := c.PostForm("apiPath")

	apis, err := service.QueryApiList(page, pageSize, apiName, apiPath)
	if err != nil {
		utils.FailMag("查询失败", c)
		return
	}

	utils.SuccessData(apis, c)
}

func (a *ApiSer) RemoveApi(c *gin.Context) {
	id := c.Query("apiId")
	err := service.RemoveApi(id)
	if err != nil {
		utils.FailMag("删除失败", c)
		return
	}
	utils.SuccessMsg("删除成功", c)
}

func (a *ApiSer) QueryApiById(c *gin.Context) {
	apiId := c.Query("id")
	data, err := service.QueryApiById(apiId)
	if err != nil {
		utils.SuccessMsg("未查询到此记录", c)
		return
	}
	utils.SuccessData(data, c)
}

func (a *ApiSer) InsertData(c *gin.Context) {
	apiName := c.Query("apiName")
	model := a.MapModel()[apiName]
	_ = c.ShouldBindJSON(&model)
	err := service.InsertData(model, apiName)
	if err != nil {
		utils.FailMag("写入数据失败", c)
		return
	}
	utils.SuccessData("写入数据成功", c)
}

func (a *ApiSer) AutoCode(path Template) {
	if path.AutoCodePath != nil {
		for _, codePath := range path.AutoCodePath {
			utils.FileInsertInfo(codePath.Path, codePath.CodeString, codePath.Delim)
		}
	}
	parse, err2 := template.ParseFiles(path.FormPath)
	if err2 != nil {
		fmt.Println("err2", err2.Error())
		return
	}
	toLower := strings.ToLower(a.api.ApiName)
	autoPath := strings.ReplaceAll(path.ToPath, "%v", toLower)
	// 文件夹
	basePath := strings.Split(autoPath, "/")
	dirPath := strings.Join(basePath[0:len(basePath)-1], "/")
	exists, errPath := utils.PathExists(dirPath)
	if errPath != nil {
		fmt.Println("errPath", errPath.Error())
		return
	}
	if !exists {
		errDir := os.MkdirAll(dirPath, 0755)
		if errDir != nil {
			fmt.Println("errDir", errDir.Error())
			return
		}
	}
	f, err3 := os.OpenFile(autoPath, os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()
	if err3 != nil {
		fmt.Println("err3", err3.Error())
		return
	}
	data := map[string]interface{}{
		"Api":            a.api,
		"ApiPackageName": strings.ToLower(a.api.ApiName),
	}
	err := parse.Execute(f, data)
	if err != nil {
		fmt.Println(err.Error(), "32222222")
	}
}

type Template struct {
	FormPath     string         `json:"form_path"`
	ToPath       string         `json:"to_path"`
	AutoCodePath []AutoCodePath `json:"auto_code_path"`
}

type AutoCodePath struct {
	Path       string `json:"path"`
	Delim      byte   `json:"delim"`
	CodeString string `json:"codeString"`
}

func (a *ApiSer) TemplatePath() []Template {
	codeData := fmt.Sprintf("auto_model.%v{},", a.api.ApiName)
	var code string
	var delim byte
	if a.api.AccessAuth {
		code = "auto_router.PrivacyRouter" + a.api.ApiName + "(PHSGroup)"
		delim = '&'
	} else {
		code = "auto_router.BaseRouter" + a.api.ApiName + "(BaseGroup)"
		delim = '#'
	}
	path := []Template{
		{
			FormPath: "./generate_code/template/router.go.tpl",
			ToPath:   "./generate_code/auto_router/%v.go",
			AutoCodePath: []AutoCodePath{
				{Path: "./initialization/router.go", Delim: delim, CodeString: code},
			},
		},
		{
			FormPath: "./generate_code/template/model.go.tpl",
			ToPath:   "./generate_code/auto_model/%v.go",
			AutoCodePath: []AutoCodePath{
				{Path: "./utils/gorm.go", Delim: '#', CodeString: codeData},
			},
		},
		{
			FormPath: "./generate_code/template/api.go.tpl",
			ToPath:   "./generate_code/auto_api/%v/%v" + a.api.ApiMethod + ".go",
		},
		{
			FormPath: "./generate_code/template/service.go.tpl",
			ToPath:   "./generate_code/auto_service/%v/%v" + a.api.ApiMethod + ".go",
		},
	}
	return path
}

func (a *ApiSer) MapModel() map[string]interface{} {
	var m auto_model.SearchList
	inter := make(map[string]interface{})
	inter["SearchList"] = m
	return inter
}
