package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

// SMTP configuration
const (
	SMTPServer = "smtp.gmail.com"
	SMTPPort   = "587"
)

var (
	SMTPUser     = os.Getenv("SMTP_USER")
	SMTPPassword = os.Getenv("SMTP_PASS")
)

func main() {
	// Check if SMTP credentials are set
	if SMTPUser == "" || SMTPPassword == "" {
		log.Fatal("SMTP credentials are not set. Please set SMTP_USER and SMTP_PASS environment variables.")
	}

	router := gin.Default()

	// Middleware to log every request
	router.Use(gin.Logger())

	// Serve static files (CSS, images, JS)
	router.Static("/static", "./static")

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/services", func(c *gin.Context) {
		c.HTML(http.StatusOK, "services.html", nil)
	})

	router.GET("/courses", func(c *gin.Context) {
		c.HTML(http.StatusOK, "courses.html", nil)
	})

	router.GET("/exclusives", func(c *gin.Context) {
		c.HTML(http.StatusOK, "exclusives.html", nil)
	})

	router.GET("/photos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "photos.html", nil)
	})

	router.GET("/support", func(c *gin.Context) {
		c.HTML(http.StatusOK, "support.html", nil)
	})

	router.GET("/videos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "videos.html", nil)
	})

	router.GET("/author", func(c *gin.Context) {
		c.HTML(http.StatusOK, "author.html", nil)
	})

	// Contact form page
	router.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	// Handle contact form submission
	router.POST("/send-message", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		message := c.PostForm("message")

		if name == "" || email == "" || message == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
			return
		}

		// Send email
		err := sendEmail(name, email, message)
		if err != nil {
			log.Println("Error sending email:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "Message sent successfully!"})
	})

	// Debug route to check static files
	router.GET("/debug-static", func(c *gin.Context) {
		files, err := os.ReadDir("./static")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read static directory", "details": err.Error()})
			return
		}

		fileNames := []string{}
		for _, file := range files {
			fileNames = append(fileNames, file.Name())
		}

		c.JSON(http.StatusOK, gin.H{"static_files": fileNames})
	})

	// Custom 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found", "path": c.Request.URL.Path})
	})

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}

// Function to send email
func sendEmail(name, email, message string) error {
	auth := smtp.PlainAuth("", SMTPUser, SMTPPassword, SMTPServer)

	to := []string{SMTPUser} // Sends email to your address
	subject := "New Contact Form Submission"
	body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage:\n%s", name, email, message)

	msg := []byte("Subject: " + subject + "\r\n" +
		"From: " + email + "\r\n" +
		"To: " + SMTPUser + "\r\n\r\n" +
		body)

	addr := SMTPServer + ":" + SMTPPort
	return smtp.SendMail(addr, auth, email, to, msg)
}


// export SMTP_USER="your-email@gmail.com"
// export SMTP_PASS="your-secure-password"
