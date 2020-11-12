package routers

import (
	peopleModule "meigo/modules/people"

	ctxExt "git.sprucetec.com/meigo/gin-context-ext"
	"github.com/gin-gonic/gin"
)

func peopleRouter(giNew *gin.Engine) {

	people := giNew.Group("/people")
	{
		people.GET("/", ctxExt.Handle(peopleModule.GetPeople))
		people.GET("/:id", ctxExt.Handle(peopleModule.GetPerson))
		people.POST("", ctxExt.Handle(peopleModule.CreatePerson))
		people.PUT("/:id", ctxExt.Handle(peopleModule.UpdatePerson))
		people.DELETE("/:id", ctxExt.Handle(peopleModule.DeletePerson))
	}
}
