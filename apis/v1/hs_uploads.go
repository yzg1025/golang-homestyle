package v1

import (
	"context"
	"gin/global"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"mime/multipart"
)

func UploadFile(c *gin.Context)  {
    file ,fileHeader,_ := c.Request.FormFile("file")
    fileSize := fileHeader.Size
    url,err := Uploads(file,fileSize)
	if err != nil {
		global.HS_LOG.Error("上传失败", zap.Any("err", err))
		utils.FailMag("上传失败", c)
		return
	}
	utils.SuccessData(url,c)
}

func Uploads(file multipart.File ,fileSize int64) (str string,err error) {
	qn := global.GCONFIG.Qiniu
	putPolicy := storage.PutPolicy{
		Scope: qn.Bucket,
	}
	mac := qbox.NewMac(qn.Accesskey,qn.Secretkey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone: &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS: false,
	}
	putExtra :=storage.PutExtra{}
	formUp := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUp.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	url := qn.Qiniuserver + ret.Key
	return url,err
}