package index

import (
	"github.com/gin-gonic/gin"
	ctxExt "github.com/kinjew/gin-context-ext"
)

/*
Index 首页
*/
func Index(c *gin.Context) {
	c.JSON(200, "Hello World!")
}

func IndexTest(c *ctxExt.Context) {
	c.JSON(200, "Hello World, c *ctxExt.Context is mostly used!")
}
