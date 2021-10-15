package common

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// 支持跨域访问
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE", "PATCH", "HEAD"},
		AllowHeaders: []string{"Authorization", "Content-Length", "X-CSRF-Token", "Accept", "Origin",
			"Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive",
			"User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma",
			"Access-Control-Allow-Origin", "X-Api-Applicationid", "Content-Disposition"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers",
			"Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "FooBar",
			"Content-Disposition"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true // 如果存在符合的origin就使用，否则用第一个AllowOrigins
		},
		AllowFiles:             true,
		AllowBrowserExtensions: true,
		MaxAge:                 2400 * time.Hour,
	})
}
