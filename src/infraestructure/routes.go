package handlers

import (
	"webhook/src/infraestructure/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {

	routes := engine.Group("pull_request")

	{
		routes.POST("process", handlers.PullRequestEvent)
	}
 //comentario
}