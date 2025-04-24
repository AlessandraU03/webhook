package application

import (
	"encoding/json"
	"fmt"
	"log"
	domain "webhook/src/domain/value_objects"
)

// Función que procesa los eventos de Workflow
func ProcessWorkflowEvent(payload []byte) int {
	var eventPayload domain.WorkflowRunEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("❌ Error al deserializar payload de workflow: %v", err)
		return 400
	}

	handleWorkflowRun(eventPayload)
	return 200
}

// Función que maneja la notificación de Workflow
func handleWorkflowRun(eventPayload domain.WorkflowRunEventPayload) {
	status := eventPayload.WorkflowRun.Status
	conclusion := eventPayload.WorkflowRun.Conclusion
	name := eventPayload.Workflow.Name
	url := eventPayload.WorkflowRun.HTMLURL
	repo := eventPayload.Repository.FullName

	message := fmt.Sprintf(
		"🚀 **Workflow Ejecutado**\n📂 **Repositorio:** %s\n🔧 **Workflow:** %s\n🟡 **Estado:** %s\n✅ **Conclusión:** %s\n🔗 **Ver detalles:** %s",
		repo, name, status, conclusion, url,
	)

	// Enviar la notificación al webhook de pruebas
	if discordWebhookPruebas == "" {
		log.Println("❌ URL del webhook de pruebas no configurada.")
		return
	}

	SendToDiscord(discordWebhookPruebas, message)
}
