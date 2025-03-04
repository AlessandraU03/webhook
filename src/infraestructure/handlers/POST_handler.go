package handlers

import (
	"webhook/src/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullRequestEvent(ctx *gin.Context) {
    eventType := ctx.GetHeader("X-GitHub-Event")

    log.Printf("Evento recibido: %s", eventType)

    if eventType != "pull_request" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Evento no soportado"})
        return
    }

    payload, err := ctx.GetRawData()
    if err != nil {
        log.Printf("Error al leer el cuerpo: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo"})
        return
    }

    statusCode := application.ProcessPullRequest(payload)

    if statusCode == 200 {
        ctx.JSON(http.StatusOK, gin.H{"status": "Evento procesado correctamente"})
    } else {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error procesando evento"})
    }
}
