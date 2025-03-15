package infraestructure

import (
	"webhook/src/infraestructure/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	routes := engine.Group("/github")

	{
		routes.POST("/webhook", handlers.WebhookHandler) // Ruta para recibir los eventos del webhook
	}
}
