package main

import (
	"log"
	"os"
	"webhook/src/infraestructure"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se encontró el archivo .env, se usarán las variables del sistema.")
	}

	// Inicializar servidor con Gin
	router := gin.Default()

	// Definir rutas
	infraestructure.Routes(router)

	// Obtener el puerto del servidor desde las variables de entorno o usar 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	log.Printf("🚀 Servidor corriendo en http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}
