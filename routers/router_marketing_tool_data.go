package routers

import (
	mtdModule "meigo/modules/marketing_tool_data"

	"github.com/gin-gonic/gin"
	ctxExt "github.com/kinjew/gin-context-ext"
)

func mtdRouter(giNew *gin.Engine) {

	mtd := giNew.Group("/mtd")
	{
		mtd.GET("/ad_query", ctxExt.Handle(mtdModule.QueryByParams))
		mtd.GET("/gst_query", ctxExt.Handle(mtdModule.GstQueryByParams))
	}
}
