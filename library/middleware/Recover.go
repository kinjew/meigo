package middleware

import (
	"meigo/library/log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

//https://blog.csdn.net/weixin_39129128/article/details/108081302
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		//加载完 defer recover，继续后续接口调用
		c.Next()
		defer func() {
			if r := recover(); r != nil {
				//打印错误堆栈信息
				log.Error("[GIN]", "panic: %v\n", r)
				debug.PrintStack()
				//封装通用json返回
				c.JSON(200, gin.H{
					"code": "4444",
					"msg":  "服务器内部错误",
				})
			}
		}()
	}
}
