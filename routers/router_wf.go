package routers

import (
	wfModule "meigo/modules/wf"

	"github.com/gin-gonic/gin"
	ctxExt "github.com/kinjew/gin-context-ext"
)

func wfRouter(giNew *gin.Engine) {

	mtd := giNew.Group("/wf")
	{
		mtd.GET("/node_change", ctxExt.Handle(wfModule.QueryByParams))
	}
}
