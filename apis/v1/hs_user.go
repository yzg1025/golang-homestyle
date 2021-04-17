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
	"go.uber.org/zap"
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
		global.HS_LOG.Warn("该用户不存在!", zap.Any("err", err))
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

func ChangePassword(c *gin.Context)  {
   var C models.ChangePassword
   _ = c.ShouldBindJSON(&C)
	if err := utils.Verify(C, utils.ChangePasswordVerify); err != nil {
		utils.FailMag(err.Error(), c)
		return
	}
	U := &models.Login{Phone: C.Phone,Password: C.Password}
	if err,_ := service.ChangePassword(U,C.NewPassword);err != nil{
		global.HS_LOG.Error("修改失败", zap.Any("err", err))
		utils.FailMag("修改失败，原密码与当前账户不符", c)
		return
	}
	utils.SuccessMsg("修改成功", c)
}

// SelectCode 查询国家code
func SelectCode(c *gin.Context)  {
	var S models.AreaCode
	_ = c.ShouldBindJSON(&S)
    err,list := service.QueryCountry()
	if err != nil {
		global.HS_LOG.Error("countrycode获取失败",zap.Any("err",err))
		utils.FailMag("countrycode获取失败",c)
	}
	utils.SuccessData(list,c)
}

// CreateCode 新增国家code
func CreateCode(c *gin.Context)  {
	var S []models.AreaCode
	_ = c.ShouldBindJSON(&S)
	for _, value := range S {
		if err:= utils.Verify(value,utils.CountryCode);err != nil{
			utils.FailMag(err.Error(),c)
			return
		}
		if isExit,country := service.IsExitCountry(value.Cname); isExit {
			global.HS_LOG.Error("重复添加",zap.Any("data",country))
			utils.FailWithDetailed(country,"重复添加",c)
			return
		}
		if err := service.CreateCountry(value);err !=nil {
			global.HS_LOG.Error("添加失败",zap.Any("err",err))
			utils.FailMag("添加失败",c)
			return
		}
	}
	utils.SuccessMsg("添加成功",c)
}

func Register(c *gin.Context)  {
	var R models.Register
	_ = c.ShouldBindJSON(&R)
	if err := utils.Verify(R,utils.Register); err != nil{
		utils.FailMag(err.Error(),c)
		return
	}
	user := &models.Login{
		NickName: R.NickName,
		Phone: R.Phone,
		Password: R.Password,
		RegisterTime:time.Now(),
		LoginMethod: "pc",
	}
    err,u := service.Register(*user)
	if err != nil {
		global.HS_LOG.Error("注册失败", zap.Any("err", err))
		utils.FailMag("注册失败",c)
		return
	}
	utils.SuccessData(u,c)
}