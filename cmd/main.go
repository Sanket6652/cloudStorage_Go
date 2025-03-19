package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/Sanket6652/cloudStorage_Go/internal/api"
	
)

func main() {
	router := gin.Default()

	// Setup API routes
	api.SetupRoutes(router)

	// Start server
	log.Println("Server running on :8080")
	router.Run(":8080")
}
