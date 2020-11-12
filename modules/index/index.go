package index

import (
	ctxExt "git.sprucetec.com/meigo/gin-context-ext"
	"github.com/gin-gonic/gin"
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
