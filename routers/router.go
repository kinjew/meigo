package routers

import (
	"net/http"

	"meigo/library/middleware"

	"github.com/gin-gonic/gin"
)

//InitRouter 路由配置
func InitRouter() *gin.Engine {
	g := gin.New()
	middlewares := []gin.HandlerFunc{}

	loadMiddle(
		g,
		middlewares...,
	)

	//加载路由配置文件
	commonRouter(g)
	//peopleRouter(g)
	exampleRouter(g)
	mtdRouter(g)
	wfRouter(g)

	return g
}

// loadMiddle loads the middleware, routes, handlers.
func loadMiddle(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middleware.
	g.Use(gin.Recovery())

	g.Use(middleware.Ginzap())

	// 性能分析
	//pprof.Register(g)

	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404！ The incorrect API route.")
	})

	return g
}
