package routers

import (
	indexModule "meigo/modules/index"

	ctxExt "git.sprucetec.com/meigo/gin-context-ext"
	"github.com/gin-gonic/gin"
)

//通用路由文件
func commonRouter(giNew *gin.Engine) {

	giNew.GET("/", indexModule.Index)
	giNew.GET("/index", ctxExt.Handle(indexModule.IndexTest))
}
