package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin/global"
	"gin/utils"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"
)

type UploadTest struct {
	OrgId            string      `json:"orgId"`
	OrgName          string      `json:"orgName"`
	AreaCode         string      `json:"areaCode"`
	AreaName         string      `json:"areaName"`
	Address          string      `json:"address"`
	GisLng           string      `json:"gisLng"`
	GisLat           string      `json:"gisLat"`
	GmtUpdate        interface{} `json:"gmtUpdate"`
	Phone            interface{} `json:"phone"`
	WorkTime         string      `json:"workTime"`
	LevelName        string      `json:"levelName"`
	OrgType          int         `json:"orgType"`
	ServiceStatus    int         `json:"serviceStatus"`
	DistanceHospital int         `json:"distanceHospital"`
	VideoUrl         interface{} `json:"videoUrl"`
	IsFree           int         `json:"isFree"`
	IsFever          int         `json:"isFever"`
	IsNeedHs         int         `json:"isNeedHs"`
	IsRed            int         `json:"isRed"`
	IsYellow         int         `json:"isYellow"`
}

type UploadFileInterFace interface {
	UploadFile(c *gin.Context)
	UploadBigFile(c *gin.Context)
}

type UploadFileStruct struct {
}

var UploadFilePart = new(UploadFileStruct)

func (upload *UploadFileStruct) UploadFile(c *gin.Context) {
	file, h, _ := c.Request.FormFile("file")
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	url, err := Uploads(h, file)
	if err != nil {
		global.HS_LOG.Error("上传失败", zap.Any("err", err))
		utils.FailMag("上传失败", c)
		return
	}
	utils.SuccessData(url, c)
}

func Uploads(h *multipart.FileHeader, f multipart.File) (string, error) {
	alyos := global.CONFIG.Oss
	extName := path.Ext(h.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
		".json": true,
		".txt":  true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}
	client, err := oss.New(alyos.Endpoint, alyos.Accesskeyid, alyos.Accesskeysecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
		return "", err
	}
	bucket, err := client.Bucket(alyos.Bucket)
	if err != nil {
		fmt.Println("获取存储空间失败")
		return "", err
	}
	fileUnixName := strconv.FormatInt(time.Now().Unix(), 10)
	saveDir := path.Join(alyos.Uploaddir, fileUnixName+extName)
	err = bucket.PutObject(saveDir, f)
	return alyos.Ossserve + saveDir, err
}

func (upload *UploadFileStruct) UploadBigFile(c *gin.Context) {
	fileh, file, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("文件获取失败")
	}
	pathDir := "./assets/upload"
	if ok, _ := utils.PathExists(pathDir); !ok {
		err := os.MkdirAll(pathDir, os.ModePerm)
		if err != nil {
			fmt.Println(fmt.Sprintf("创建目录失败%s", err))
			return
		}
	}
	fileDest := path.Join(pathDir, file.Filename)
	errorSave := c.SaveUploadedFile(file, fileDest)
	if errorSave != nil {
		fmt.Println("保存文件失败")
		return
	}
	fs, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", pathDir, file.Filename))
	if err != nil {
		fmt.Println(err)
	}
	content := string(fs)
	var uploadModel []UploadTest
	err = json.Unmarshal([]byte(content), &uploadModel)
	errRemove := os.RemoveAll(pathDir)
	if errRemove != nil {
		fmt.Println("删除文件失败")
		return
	}
	fileUrl, err := Uploads(file, fileh)
	if err != nil {
		fmt.Println("上传文件失败", err)
		return
	}
	maps := map[string]interface{}{
		"analyzeData": uploadModel,
		"url":         fileUrl,
	}
	utils.SuccessData(maps, c)
}
