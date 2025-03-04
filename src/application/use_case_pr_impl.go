package application

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	domain "webhook/src/domain/value_objects"
)

const discordWebhookURL = "https://discord.com/api/webhooks/1346298566649315349/XfWtXCowZsmu9ek4bR_u2XCQgp5NxNAZ_TSTIN1PL5hJKrX_t7XDnZiFIb1dXEcRcVg3"

func sendToDiscord(message string) {
    payload := map[string]string{"content": message}
    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error al serializar JSON: %v", err)
        return
    }

    resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer(payloadBytes))
    if err != nil {
        log.Printf("Error al enviar mensaje a Discord: %v", err)
        return
    }
    defer resp.Body.Close()

    log.Printf("Respuesta de Discord: %v", resp.Status)
}


func ProcessPullRequest(payload []byte) int {
    var eventPayload domain.PullRequestEventPayload

    if err := json.Unmarshal(payload, &eventPayload); err != nil {
        log.Printf("Error al deserializar payload: %v", err)
        return 400
    }

    if eventPayload.Action != "opened" && eventPayload.Action != "edited" {
        log.Printf("Evento no compatible: %s", eventPayload.Action)
        return 400
    }

    user := eventPayload.PullRequest.User.Login
    title := eventPayload.PullRequest.Title
    url := eventPayload.PullRequest.URL
    comment := eventPayload.PullRequest.Body // Obtiene el comentario

    message := "**Nuevo Pull Request**\n" +
        "👤 **Usuario:** " + user + "\n" +
        "📌 **Título:** " + title + "\n" +
        "💬 **Comentario:** " + comment + "\n" + // Agregar comentario
        "🔗 **URL:** " + url

    log.Println("Enviando mensaje a Discord...")
    sendToDiscord(message)

    return 200
}


