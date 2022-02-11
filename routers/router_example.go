package routers

import (
	"meigo/gqlgen-todos/graph"
	"meigo/gqlgen-todos/graph/generated"
	exampleModule "meigo/modules/example"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/99designs/gqlgen/graphql/handler/extension"

	"github.com/99designs/gqlgen/graphql/playground"

	ctxExt "github.com/kinjew/gin-context-ext"

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
		example.POST("/query", graphqlHandler())
		example.GET("/", playgroundHandler())
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//// Using schema directives to implement permission checks,example
	//c := generated.Config{Resolvers: &graph.Resolver{}}
	//c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
	//	fmt.Println(role, "input role")
	//	/*
	//		if !getCurrentUser(ctx).HasRole(role) {
	//			// block calling the next resolver
	//			return nil, fmt.Errorf("Access denied")
	//		}
	//	*/
	//	//return nil, fmt.Errorf("Access denied")
	//	// or let it pass through
	//	return next(ctx)
	//}
	//h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	h.Use(extension.FixedComplexityLimit(5)) // This line is key,any query with complexity greater than 5 is rejected by the API.
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/example/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
