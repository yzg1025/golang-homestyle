package v1

import (
	"errors"
	"fmt"
	"gin/global"
	"gin/utils"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadFile(c *gin.Context) {
	file, h, _ := c.Request.FormFile("file")
	defer file.Close()

	url, err := Uploads(h, file)
	if err != nil {
		global.HS_LOG.Error("上传失败", zap.Any("err", err))
		utils.FailMag("上传失败", c)
		return
	}
	utils.SuccessData(url, c)
}

func Uploads(h *multipart.FileHeader, f multipart.File) (string, error) {
	alyos := global.GCONFIG.Oss
	extName := path.Ext(h.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
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
