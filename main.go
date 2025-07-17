package main

import (
	"log"
	"tryg/config"
	"tryg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()
	routes.RegisterRoutes(router)

	log.Println("Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}