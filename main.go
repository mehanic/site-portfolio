// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/smtp"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	_ "github.com/lib/pq" // PostgreSQL driver
// )

// // SMTP Configuration
// const (
// 	SMTPServer = "smtp.gmail.com"
// 	SMTPPort   = "587"
// )

// var (
// 	SMTPUser     = os.Getenv("SMTP_USER")
// 	SMTPPassword = os.Getenv("SMTP_PASS")
// 	DBConnStr    = os.Getenv("DATABASE_URL") // Database connection string
// )

// var db *sql.DB // Database connection

// func main() {

// 	loadEnv()
// 	// Check if SMTP credentials are set
// 	if SMTPUser == "" || SMTPPassword == "" {
// 		log.Fatal("SMTP credentials are not set. Please set SMTP_USER and SMTP_PASS environment variables.")
// 	}

// 	// Connect to PostgreSQL
// 	var err error
// 	db, err = sql.Open("postgres", DBConnStr)
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}
// 	defer db.Close()

// 	// Test database connection
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Database connection failed:", err)
// 	}
// 	fmt.Println("Connected to PostgreSQL database!")

// 	router := gin.Default()

// 	// Set trusted proxies to avoid the warning (adjust this based on your proxy configuration)
// 	router.SetTrustedProxies([]string{"127.0.0.1"}) // Replace with your trusted proxy IPs if needed

// 	// Middleware to log every request
// 	router.Use(gin.Logger())

// 	// Serve static files (CSS, images, JS)
// 	router.Static("/static", "./static")

// 	// Load HTML templates
// 	router.LoadHTMLGlob("templates/*")

// 	// Define routes
// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", nil)
// 	})

// 	router.GET("/services", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "services.html", nil)
// 	})

// 	router.GET("/courses", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "courses.html", nil)
// 	})

// 	router.GET("/exclusives", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "exclusives.html", nil)
// 	})

// 	router.GET("/photos", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "photos.html", nil)
// 	})

// 	router.GET("/support", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "support.html", nil)
// 	})

// 	router.GET("/videos", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "videos.html", nil)
// 	})

// 	router.GET("/author", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "author.html", nil)
// 	})

// 	// Contact form page
// 	router.GET("/contact", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "contact.html", nil)
// 	})

// 	// Handle contact form submission
// 	router.POST("/send-message", func(c *gin.Context) {
// 		name := c.PostForm("name")
// 		email := c.PostForm("email")
// 		message := c.PostForm("message")

// 		if name == "" || email == "" || message == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
// 			return
// 		}

// 		// Store message in PostgreSQL
// 		err := saveMessageToDB(name, email, message)
// 		if err != nil {
// 			log.Println("Error saving message to database:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
// 			return
// 		}

// 		// Send email notification
// 		err = sendEmail(name, email, message)
// 		if err != nil {
// 			log.Println("Error sending email:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"success": "Message sent and saved successfully!"})
// 	})

// 	// Debug route to check static files
// 	router.GET("/debug-static", func(c *gin.Context) {
// 		files, err := os.ReadDir("./static")
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read static directory", "details": err.Error()})
// 			return
// 		}

// 		fileNames := []string{}
// 		for _, file := range files {
// 			fileNames = append(fileNames, file.Name())
// 		}

// 		c.JSON(http.StatusOK, gin.H{"static_files": fileNames})
// 	})

// 	// Custom 404 handler
// 	router.NoRoute(func(c *gin.Context) {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found", "path": c.Request.URL.Path})
// 	})

// 	// Start the server
// 	fmt.Println("Server running on http://localhost:8080")
// 	router.Run(":8080")
// }

// // Load environment variables from .env file

// // Load environment variables from .env file
// func loadEnv() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	SMTPUser = os.Getenv("SMTP_USER")
// 	SMTPPassword = os.Getenv("SMTP_PASS")
// 	DBConnStr = os.Getenv("DATABASE_URL")
// }

// // Handle contact form submission
// func handleSendMessage(c *gin.Context) {
// 	name := c.PostForm("name")
// 	email := c.PostForm("email")
// 	message := c.PostForm("message")

// 	if name == "" || email == "" || message == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
// 		return
// 	}

// 	if err := saveMessageToDB(name, email, message); err != nil {
// 		log.Println("Error saving message to database:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
// 		return
// 	}

// 	if err := sendEmail(name, email, message); err != nil {
// 		log.Println("Error sending email:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"success": "Message sent and saved successfully!"})
// }

// // Function to save message to PostgreSQL
// func saveMessageToDB(name, email, message string) error {
// 	_, err := db.Exec("INSERT INTO contact_messages (name, email, message) VALUES ($1, $2, $3)", name, email, message)
// 	return err
// }

// // Function to send email
// func sendEmail(name, email, message string) error {
// 	auth := smtp.PlainAuth("", SMTPUser, SMTPPassword, SMTPServer)

// 	to := []string{SMTPUser} // Sends email to your address
// 	subject := "New Contact Form Submission"
// 	body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage:\n%s", name, email, message)

// 	msg := []byte("Subject: " + subject + "\r\n" +
// 		"From: " + email + "\r\n" +
// 		"To: " + SMTPUser + "\r\n\r\n" +
// 		body)

// 	addr := SMTPServer + ":" + SMTPPort
// 	return smtp.SendMail(addr, auth, email, to, msg)
// }
// func logError(context string, err error) {
// 	if err != nil {
// 		log.Printf("[ERROR] %s: %v", context, err)
// 	}
// }


// func setupGracefulShutdown(router *gin.Engine) {
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

// 	go func() {
// 		<-quit
// 		fmt.Println("\nShutting down server...")
// 		db.Close()
// 		os.Exit(0)
// 	}()
// }
// // export SMTP_USER="your-email@gmail.com"
// // export SMTP_PASS="your-secure-password"
// //export DATABASE_URL="postgres://mehanic:new_password@localhost:5432/portfolio?sslmode=require"
