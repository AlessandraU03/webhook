package handlers

import (
	"log"
	"net/http"
	"webhook/src/application"

	"github.com/gin-gonic/gin"
)

// WebhookHandler maneja los eventos de GitHub, como 'pull_request' y 'workflow_run'
func WebhookHandler(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")
	signature := ctx.GetHeader("X-Hub-Signature-256")

	log.Println(signature)
	log.Printf("📩 Webhook recibido: \nEvento=%s, \nDeliveryID=%s", eventType, deliveryID)

	payload, err := ctx.GetRawData()
	if err != nil {
		log.Printf("❌ Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	var statusCode int

	switch eventType {
	case "pull_request":
		statusCode = application.ProcessPullRequest(payload)
	case "workflow_run":
		statusCode = application.ProcessWorkflowEvent(payload)
	default:
		log.Printf("⚠️ Evento no compatible: %s", eventType)
		statusCode = 400
	}

	ctx.JSON(statusCode, gin.H{"status": "Evento procesado"})
}
