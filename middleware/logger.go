package middleware

import (
	"fmt"
	"go-skeleton/utils/log"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		path := c.Request.RequestURI
		method := c.Request.Method
		clientIp := c.ClientIP()
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		userAgent := c.Request.UserAgent()
		log.Infof("msg in << Path: %v, Method: %v, HostName: %v, Ip: %v, Agent: %v", path, method, hostName, clientIp, userAgent)

		c.Next()

		stopTime := time.Since(startTime)
		spendTime := int(math.Ceil(float64(stopTime.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}

		// logger := log.GetLogger()
		// ety := logger.WithFields(logrus.Fields{
		// 	"HostName":   hostName,
		// 	"StatusCode": statusCode,
		// 	"SpendTime":  spendTime,
		// 	"Ip":         clientIp,
		// 	"Method":     method,
		// 	"Path":       path,
		// 	"DataSize":   dataSize,
		// 	"Agent":      userAgent,
		// })
		logStr := fmt.Sprintf("msg out >> Path: %v, Method: %v, HostName: %v, StatusCode: %v, SpendTime: %v, Ip: %v, DataSize: %v, Agent: %v",
			path, method, hostName, statusCode, spendTime, clientIp, dataSize, userAgent)

		if len(c.Errors) > 0 {
			log.Errorf(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			log.Errorf(logStr)
		} else if statusCode >= 400 {
			log.Warnf(logStr)
		} else {
			log.Infof(logStr)
		}
	}
}
