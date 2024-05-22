package middleware

import (
	"github.com/spf13/viper"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger предоставляет дополнительные логи с помощью библиотеки logrus.
// Данный middleware можно отключить в файлах конфигурации, установив "logrus_middleware" в значение 'false'
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		latency := time.Since(startTime)

		statusCode := c.Writer.Status()

		entry := logrus.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency":     latency,
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.String())
		} else {
			entry.Print("Request completed")
		}
	}
}

// CheckCfg мапит файл конфига и в случае logrus_middleware = true включает middleware logger и устанавливает уровень
// логгирования для logrus
func CheckCfg(router *gin.Engine) {
	if viper.GetBool("logrus_middleware") {
		router.Use(Logger())
		level, err := logrus.ParseLevel(viper.GetString("log_level"))
		if err != nil {
			logrus.Fatalf("invalid log level: %s", err.Error())
		}
		logrus.SetLevel(level)
	}
}
