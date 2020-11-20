package routers

import (
	indexModule "meigo/modules/index"

	"github.com/gin-gonic/gin"
	ctxExt "github.com/kinjew/gin-context-ext"
)

//通用路由文件
func commonRouter(giNew *gin.Engine) {

	giNew.GET("/", indexModule.Index)
	giNew.GET("/index", ctxExt.Handle(indexModule.IndexTest))
}
