package infraestructure

import (
	"webhook/src/infraestructure/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	routes := engine.Group("webhook")
	{
		routes.POST("pull_request/process", handlers.WebhookHandler)
		routes.POST("workflow/process", handlers.WebhookHandler)
	}
}
