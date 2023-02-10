package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
	"net/http"
	"path/filepath"
	"tcpSimulator/common"
)

func Http(host string, relativePath string, certPath string) {
	r := gin.Default()
	r.Use(common.Cors())

	//　支持跨域
	r.Any(relativePath, handler)

	if certPath != "" {
		if err := common.Generate(certPath); err != nil {
			logrus.Fatalln(err)
		}
		r.Use(tlsHandler())
		// 支持https
		logrus.Fatalln(
			r.RunTLS(host,
				filepath.Join(certPath, "server.pem"),
				filepath.Join(certPath, "server.key"),
			),
		) // 端口被占用的时候提示错误
	}

	// 任意http的接口类型
	logrus.Fatalln(r.Run(host)) // 端口被占用的时候提示错误
}

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"state":   http.StatusText(http.StatusOK),
		"path":    c.Param("any"),
		"time":    common.TimeNow(),
		"address": common.GetLocalIP(),
	})
}

func tlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			FrameDeny: true,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		// If there was an error, do not continue.
		if err != nil {
			c.Abort()
			return
		}
		// Avoid header rewrite if response is a redirection
		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}
		c.Next()
	}
}
