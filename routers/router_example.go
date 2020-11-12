package routers

import (
	exampleModule "meigo/modules/example"

	ctxExt "git.sprucetec.com/meigo/gin-context-ext"

	"github.com/gin-gonic/gin"
)

//实例路由，便于后续路由按文件书写
func exampleRouter(giNew *gin.Engine) {
	example := giNew.Group("/example")
	{
		example.POST("/upload-single", ctxExt.Handle(exampleModule.UploadSingle))
		example.POST("/upload-multiple", ctxExt.Handle(exampleModule.UploadMultiple))
		example.POST("/handle-go", ctxExt.Handle(exampleModule.HandleGo))
		example.GET("/valid-bookable", ctxExt.Handle(exampleModule.ValidBookable))
		example.GET("/redis", ctxExt.Handle(exampleModule.Redis))
	}
}
