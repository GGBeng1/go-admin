package v1

import (
	"fmt"
	"hello/helper"
	"hello/model/request"
	"net/http"

	"github.com/dchest/captcha"
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
		// tokenNext(c,L)
		// fmt.Println()
		helper.SuccessWithData(L, c)
	} else {
		helper.FailedWithMsg("验证码错误", c)
	}
}
