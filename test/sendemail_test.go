package test

import (
	"crypto/tls"
	"fmt"
	"gin/global"
	"gin/utils"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"math/rand"
	"net/smtp"
	"strings"
	"testing"
	"time"
)

func TestName1(t *testing.T) {
	e := email.NewEmail()
	mailUserName := "yzg862689234@163.com" //邮箱账号
	mailPassword := "WFXJVMSBAPNHRJMC"     //邮箱授权码 //发送的验证码
	Subject := "提供验证码"                     //发送的主题

	e.From = "Demo <yzg862689234@163.com>"
	e.To = []string{"862689234@qq.com"}
	e.Subject = Subject
	sendTempl := `
      <div style="background:rgba(2,132,199,0.1);border-radius:5px;padding: 12px">
        <img style="width:30px;height:30px;border-radius:15px" src="https://wx.qlogo.cn/mmhead/Q3auHgzwzM5WcGYUnaIqVFaqfaEoy3ZCt5NibmMBaQjzQaNOQaybeiag/0"/>
        <div>您的验证码:%v,验证码有效期5分钟</div>
        <div style="width:100vw;height:1px;background:#eee"></div>   
        <div>如非本人操作，请忽略此邮件</div>
      </div>
    `
	html := fmt.Sprintf(sendTempl, utils.GetRandom())
	e.HTML = []byte(html)
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", mailUserName, mailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		global.HS_LOG.Warn("发送验证码失败!", zap.Any("err", err))
	}
}

func TestName2(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	t.Fatal(fmt.Sprintf("%06v", rnd.Int31n(100000)))
}

func TestPath(t *testing.T) {
	path := "./generate_code/auto_api/indicator/indicatorQueryGet.go"
	basePath := strings.Split(path, "/")
	filePath := basePath[len(basePath)-1]
	dirPathArr := basePath[0 : len(basePath)-1]
	dirPath := strings.Join(dirPathArr, "/")
	t.Fatal("32323", filePath, dirPath)
}

// 判断所给路径是否为文件夹

func TestIsDir(t *testing.T) {
	exists, err := utils.PathExists("./generate_code/auto_api")
	if err != nil {
		return
	}
	t.Log("323", exists)
}
