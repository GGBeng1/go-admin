package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 400
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func SuccessWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func SuccessWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func FailedWithMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}
