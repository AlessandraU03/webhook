package application

import (
	"encoding/json"
	"fmt"
	"log"
	domain "webhook/src/domain/value_objects"
)

// Función que procesa los eventos de Pull Request
func ProcessPullRequest(payload []byte) int {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("❌ Error al deserializar payload: %v", err)
		return 400
	}

	handlePullRequest(eventPayload)
	return 200
}

// Función que maneja la notificación de Pull Request
func handlePullRequest(eventPayload domain.PullRequestEventPayload) {
	user := eventPayload.PullRequest.User.Login
	title := eventPayload.PullRequest.Title
	url := eventPayload.PullRequest.URL
	action := eventPayload.Action

	message := fmt.Sprintf(
		"📢 **Pull Request %s**\n👤 **Usuario:** %s\n📌 **Título:** %s\n🔗 **URL:** %s",
		action, user, title, url,
	)

	// Enviar la notificación al webhook de desarrollo
	if discordWebhookDesarrollo == "" {
		log.Println("❌ URL del webhook de desarrollo no configurada.")
		return
	}
    

	SendToDiscord(discordWebhookDesarrollo, message)

	// Si el PR ha sido fusionado, también se envía una notificación adicional
	if action == "closed" {
		SendToDiscord(discordWebhookDesarrollo, "✅ **El PR ha sido fusionado exitosamente!**")
	}
}
