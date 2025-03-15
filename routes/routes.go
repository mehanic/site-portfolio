package routes

import (
	"site-portfolio/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Load HTML templates
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Define routes
	router.GET("/", handlers.HandleHome)
	router.GET("/contact", handlers.HandleContact)

	router.GET("/services", handlers.HandleServices)
	router.GET("/exclusives", handlers.HandleExclusives)
	router.GET("/support", handlers.HandleSupport)

	// If you want a `/courses` route, add this:
	router.GET("/courses", handlers.HandleCourses)

	// Handle contact form submissions
	router.POST("/send-message", handlers.HandleSendMessage)

	// 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Page not found"})
	})

	router.GET("/photos", handlers.HandlePhotos)
	router.GET("/videos", handlers.HandleVideos)
	router.GET("/author", handlers.HandleAuthor)
}
