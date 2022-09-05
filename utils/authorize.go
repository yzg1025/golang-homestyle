package utils

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"gin/global"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"io"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"
)

// GetRandom 生成6位随机数
func GetRandom() string {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

// SendEmail 发送邮箱验证码
func SendEmail(toAccount, code string) error {
	e := email.NewEmail()
	mailUserName := "yzg862689234@163.com" //邮箱账号
	mailPassword := "WFXJVMSBAPNHRJMC"     //邮箱授权码 //发送的验证码
	Subject := "提供验证码"                     //发送的主题

	e.From = "Demo <yzg862689234@163.com>"
	e.To = []string{toAccount}
	e.Subject = Subject
	sendTempl := `
      <div style="background:rgba(2,132,199,0.1);border-radius:5px;padding: 12px">
        <img style="width:30px;height:30px;border-radius:15px" src="https://wx.qlogo.cn/mmhead/Q3auHgzwzM5WcGYUnaIqVFaqfaEoy3ZCt5NibmMBaQjzQaNOQaybeiag/0"/>
        <div>您的验证码:%v,验证码有效期5分钟</div>
        <div style="width:100vw;height:1px;background:#eee"></div>   
        <div>如非本人操作，请忽略此邮件</div>
      </div>
    `
	html := fmt.Sprintf(sendTempl, code)
	e.HTML = []byte(html)
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", mailUserName, mailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		global.HS_LOG.Warn("发送验证码失败!", zap.Any("err", err))
		return err
	}
	return nil
}

// FileInsertInfo 代码插入
func FileInsertInfo(filepath, codeData string, delim byte) {
	temp := strings.ReplaceAll(filepath, ".go", ".tmp")
	file, err := os.OpenFile(filepath, os.O_RDWR, 0544)
	if err != nil {
		fmt.Printf("File open failed! err: %v\n", err)
		return
	}
	reader := bufio.NewReader(file)
	// 新建临时文件
	tempFile, err := os.OpenFile(temp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
	defer tempFile.Close()
	if err != nil {
		fmt.Printf("Temp create failed! err: %v\n", err)
		return
	}
	writer := bufio.NewWriter(tempFile)
	_ = writer.Flush()
	// 将原文件写入临时文件
	line, err := reader.ReadString(delim)
	if err != nil {
		fmt.Printf("Read file failed! err: %v\n", err)
		return
	}
	// 写入临时文件
	w := fmt.Sprintf("\n%v\n", line)
	_, _ = writer.WriteString(w)
	_ = writer.Flush()
	// 写入要插入的内容
	_, _ = tempFile.WriteString(codeData)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("File raed failed! err: %v\n", err)
			return
		}
		_, _ = writer.WriteString(line)
	}
	_ = writer.Flush()
	err = os.Rename(temp, filepath)
	if err != nil {
		fmt.Printf("Rename file raed failed! err: %v\n", err)
		return
	}
}
