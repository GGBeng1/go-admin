package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logFilePath := "./"
	logFileName := "log.log"

	//文件路径
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		// 执行时间
		allTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUrl := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求ip
		clientIP := c.ClientIP()

		// 换一下日期格式
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		// 换成json格式
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		// 日志变成json
		logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"all_time":    allTime,
			"client_ip":   clientIP,
			"req_method":  reqMethod,
			"req_url":     reqUrl,
		}).Info()

		// 日志是字符串
		// logger.Infof("| %3d | %13v | %15s | %s | %s |",
		//  statusCode,
		//  allTime,
		//  clientIP,
		//  reqMethod,
		//  reqUrl,
		// )
	}
}
