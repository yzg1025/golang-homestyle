package v1

import (
	"fmt"
	"gin/global"
	"gin/middleware"
	"gin/models"
	"gin/service"
	"gin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Login(c *gin.Context)  {
	var L models.Login
	_ = c.ShouldBindJSON(&L)
    if err:= utils.Verify(L,utils.LoginVerify);err != nil{
   	  utils.FailMag(err.Error(),c)
	  return
    }
    err,u:=service.Login(L.Phone,L.Password)
	if err != nil {
		utils.FailMag("该用户不存在",c)
		return
	}
	token := tokenSend(L)
    data := map[string]interface{}{
    	"data":u,
    	"token":token,
	}
	utils.SuccessData(data,c)
}

func tokenSend(user models.Login) string {
	j := &middleware.JWT{
		SigningKey: []byte(global.GCONFIG.JWT.Signingkey),
	}
	claims := middleware.MyClaims{
		ID: user.ID,
		NickName: user.NickName,
		BufferTime: 60 * 60 * 24,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + 60*60*24*7,
			Issuer: global.GCONFIG.JWT.Signingkey,
		},
	}
	token,err := j.CreateToken(claims)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return token
}

// 发送验证码
func SendMsgCode(c *gin.Context)  {
	var L models.Login
	_ = c.ShouldBindJSON(&L)
	if err:= utils.Verify(L,utils.SendCode);err != nil{
		utils.FailMag(err.Error(),c)
		return
	}

	var jh = global.GCONFIG
	//初始化参数
	param := url.Values{}
	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("mobile", L.Phone) //接受短信的用户手机号码
	param.Set("tpl_id", strconv.Itoa(jh.JUHE.Tplid)) //您申请的短信模板ID，根据实际情况修改
	param.Set("tpl_value", jh.JUHE.Tplvalue) //您设置的模板变量，根据实际情况
	param.Set("key", jh.JUHE.Apikey)  //应用APPKEY(应用详细页查询)

	data,err:= Post(jh.JUHE.Juheurl,param)
	if err != nil {
		fmt.Println(err.Error())
	}
    utils.SuccessData(data,c)
}

func Post(apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

