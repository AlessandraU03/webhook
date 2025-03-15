package application

import (
	"encoding/json"
	"fmt"
	"log"
	domain "webhook/src/domain/value_objects"
)

// ProcessPullRequest maneja las notificaciones de Pull Requests
func ProcessPullRequest(payload []byte) int {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("❌ Error al deserializar payload: %v", err)
		return 400
	}

	handlePullRequest(eventPayload)
	return 200
}

func handlePullRequest(eventPayload domain.PullRequestEventPayload) {
	user := eventPayload.PullRequest.User.Login
	title := eventPayload.PullRequest.Title
	url := eventPayload.PullRequest.URL
	action := eventPayload.Action

	message := fmt.Sprintf(
		"📢 **Pull Request %s**\n👤 **Usuario:** %s\n📌 **Título:** %s\n🔗 **URL:** %s",
		action, user, title, url,
	)

	sendToDiscord(discordWebhookDesarrollo, message)

	// Si el PR se fusionó, también lo notificamos
	if action == "closed" {
		sendToDiscord(discordWebhookDesarrollo, "✅ **El PR ha sido fusionado exitosamente!**")
	}
}
