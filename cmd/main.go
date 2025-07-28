// package main

// import (
// 	"fmt"
// 	"log"

// 	"site-portfolio/config"
// 	"site-portfolio/routes"
// 	"site-portfolio/utils"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Load environment variables
// 	config.LoadEnv()

// 	// Connect to database
// 	config.InitDB()
// 	defer config.CloseDB()

// 	// Set up router
// 	router := gin.Default()
// 	routes.SetupRoutes(router)

// 	// Handle graceful shutdown
// 	utils.SetupGracefulShutdown(router)

// 	// fmt.Println("Server running on http://localhost:40000")
// 	// log.Fatal(router.Run(":40000"))

// 	host := "0.0.0.0"
// 	port := "40001"

// 	// Define certificate paths
// 	// certFile := "/home/mehanic/site-portfolio/certificate/cert.pem"
// 	// keyFile := "/home/mehanic/site-portfolio/certificate/key.pem"

// 	fmt.Printf("Server running on http://%s:%s\n", host, port)
// 	log.Fatal(router.Run(host + ":" + port)) // Bind to 0.0.0.0:40000

// }

package main

import (
	"fmt"
	"log"
	"os"

	"site-portfolio/config"
	"site-portfolio/handlers"
	"site-portfolio/routes"
	"site-portfolio/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	
    err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	// проверка:
	log.Println("EMAIL:", os.Getenv("SMTP_EMAIL"))
	log.Println("PASS:", os.Getenv("SMTP_PASSWORD"))
	log.Println("TO:", os.Getenv("SMTP_RECIPIENT"))


	config.LoadEnv()
	config.InitDB()
	defer config.CloseDB()

	router := gin.Default()
	routes.SetupRoutes(router)

	utils.SetupGracefulShutdown(router)
	handlers.SetupContactRoutes(router)
	host := "127.0.0.1"
	port := "8080"

	fmt.Printf("Server running on http://%s:%s\n", host, port)
	log.Fatal(router.Run(host + ":" + port))
}
