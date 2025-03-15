package main

import (
	"fmt"
	"log"

	"site-portfolio/config"
	"site-portfolio/routes"
	"site-portfolio/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	config.InitDB()
	defer config.CloseDB()

	// Set up router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Handle graceful shutdown
	utils.SetupGracefulShutdown(router)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
