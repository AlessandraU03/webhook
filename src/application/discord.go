package application

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)


var (
	discordWebhookDesarrollo = os.Getenv("DISCORD_WEBHOOK_DESARROLLO")
	discordWebhookPruebas    = os.Getenv("DISCORD_WEBHOOK_PRUEBAS")
	discordWebhookGeneral    = os.Getenv("DISCORD_WEBHOOK_GENERAL")
)


// Función para enviar mensaje a Discord
func SendToDiscord(webhookURL, message string) {
	if webhookURL == "" {
		log.Println("❌ No se ha configurado la URL del webhook.")
		return
	}

	payload := map[string]string{"content": message}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("❌ Error al serializar JSON: %v", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("❌ Error al enviar mensaje a Discord: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("✅ Respuesta de Discord: %v", resp.Status)
}
