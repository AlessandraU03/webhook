package handlers

import (
	"webhook/src/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullRequestEvent(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")
	signature := ctx.GetHeader("X-Hub-Signature-256")

	log.Println("Firma del webhook:", signature)
	log.Printf("Webhook recibido: \nEvento=%s, \nDeliveryID=%s", eventType, deliveryID)

	payload, err := ctx.GetRawData()
	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	var statusCode int

	switch eventType {
	case "pull_request":
		statusCode = application.ProcessPullRequest(payload)
	case "issue_comment":
		statusCode = application.ProcessCommentEvent(payload)
	default:
		log.Printf("Evento no soportado: %s", eventType)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Evento no soportado"})
		return
	}

	if statusCode == 200 {
		ctx.JSON(http.StatusOK, gin.H{"status": "Evento procesado correctamente"})
	} else {
		log.Printf("Error al procesar el evento: %s", eventType)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el evento"})
	}
}
