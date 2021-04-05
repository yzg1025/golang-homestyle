package v1

import (
	"fmt"
	"gin/global"
	"gin/middleware"
	"gin/models"
	"gin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context)  {
	var L models.Login
	_ = c.ShouldBindJSON(&L)
    if err:= utils.Verify(L,utils.LoginVerify);err != nil{
   	  utils.FailMag(err.Error(),c)
	  return
    }
	tokenSend(c,L)
}

func tokenSend(c *gin.Context,user models.Login)  {
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
			Issuer: "homestyle",
		},
	}
	token,err := j.CreateToken(claims)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.SuccessData(token,c)
}

