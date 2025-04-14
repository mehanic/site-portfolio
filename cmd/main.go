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

	// fmt.Println("Server running on http://localhost:40000")
	// log.Fatal(router.Run(":40000"))

	host := "0.0.0.0"
	port := "40001"

	// Define certificate paths
	// certFile := "/home/mehanic/site-portfolio/certificate/cert.pem"
	// keyFile := "/home/mehanic/site-portfolio/certificate/key.pem"

	fmt.Printf("Server running on http://%s:%s\n", host, port)
	log.Fatal(router.Run(host + ":" + port)) // Bind to 0.0.0.0:40000
}
