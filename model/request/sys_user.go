package request

import "github.com/dgrijalva/jwt-go"

type RigisterStruct struct {
	Username string `json:"username" binding:"required" `
	Password string `json:"password" binding:"required" `
}
type LoginStruct struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captchaId" binding:"required"`
}
type CaptchaResponse struct {
	Path      string `json:"path" `
	CaptchaId string `json:"captchaId"`
}

// Custom claims structure
type CustomClaims struct {
	Username string
	jwt.StandardClaims
}
