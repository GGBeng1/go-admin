package v1

import (
	"fmt"
	"hello/helper"
	"hello/middleware"
	"hello/model"
	"hello/model/request"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var R request.RigisterStruct
	if err := c.ShouldBindJSON(&R); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(R)
	// c.JSON(http.StatusOK, gin.H{"status": R})
	helper.SuccessWithMsg("注册成功", c)
}
func Login(c *gin.Context) {
	var L request.LoginStruct
	_ = c.ShouldBindJSON(&L)
	if captcha.VerifyString(L.CaptchaId, L.Captcha) {
		tokenNext(c, L)
		// fmt.Println()
		// helper.SuccessWithData(L, c)
	} else {
		helper.FailedWithMsg("验证码错误", c)
	}
}
func tokenNext(c *gin.Context, user request.LoginStruct) {
	j := &middleware.JWT{
		[]byte(model.JwtSignStruct), // 唯一签名
	}
	clams := request.CustomClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "ggbeng",                              //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		helper.FailedWithMsg("获取token失败", c)
	} else {
		helper.SuccessWithData(token, c)
	}
}

func TestToken(c *gin.Context) {
	helper.SuccessWithMsg("权限通过", c)
}
