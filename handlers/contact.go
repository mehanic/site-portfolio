package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"site-portfolio/email"   
	"site-portfolio/models"
)

// HandleSendMessage processes the contact form submission
func HandleSendMessage(c *gin.Context) {
	name := c.PostForm("name")
	senderEmail := c.PostForm("email")
	message := c.PostForm("message")

	if name == "" || senderEmail == "" || message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	if err := models.SaveMessage(name, senderEmail, message); err != nil {
		log.Println("Error saving message to database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}

	if err := email.SendEmail(name, senderEmail, message); err != nil {  
		log.Println("Error sending email:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Message sent and saved successfully!"})
}
