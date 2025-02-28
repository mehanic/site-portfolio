// package main

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// var pageData = map[string]interface{}{
// 	"index": gin.H{
// 		"Title": "DevOps Compass | Guided IT Solutions by Docker Captain",
// 	},
// 	"services": gin.H{
// 		"Title": "Services | DevOps Compass",
// 		"Services": []string{
// 			"DevOps Transformation",
// 			"Cloud Solutions",
// 			"Security Audits and Compliance",
// 			"Docker & Containerization Expertise",
// 			"Infrastructure as Code (IaC) Implementation",
// 			"Kubernetes Orchestration",
// 		},
// 		"Email": "contact@devopscompass.com",
// 	},
// 	// Add more pages like "exclusives", "photos", etc., if needed.
// }

// func main() {
// 	router := gin.Default()

// 	// Serve static files (CSS, images, JS)
// 	router.Static("/static", "./static")

// 	// Load HTML templates
// 	router.LoadHTMLGlob("templates/*.html")

// 	// Route for Home Page
// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "layout.html", pageData["index"])
// 	})

// 	// Route for Services Page
// 	router.GET("/services", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "layout.html", pageData["services"])
// 	})

// 	// Route for Courses Page
// 	router.GET("/courses", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "layout.html", pageData["courses"])
// 	})

// 	// Generic route for other pages (exclusives, photos, etc.)
// 	pages := []string{"exclusives", "photos", "support", "author"}
// 	for _, page := range pages {
// 		page := page // Fix closure issue
// 		router.GET("/"+page, func(c *gin.Context) {
// 			data, exists := pageData[page]
// 			if !exists {
// 				data = gin.H{"Title": page} // Default title if no data is defined
// 			}
// 			c.HTML(http.StatusOK, "layout.html", data)
// 		})
// 	}

//		// Start the server
//		router.Run(":8080")
//	}
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middleware to log every request
	router.Use(gin.Logger())

	// Serve static files (CSS, images, JS)
	router.Static("/static", "./static")

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

	// ✅ Corrected: No need to check error for LoadHTMLGlob
	router.LoadHTMLGlob("templates/*.html")

	// Route for Home Page
	router.GET("/", func(c *gin.Context) {
		renderPage(c, "index")
	})

	// Route for Services Page
	router.GET("/services", func(c *gin.Context) {
		renderPage(c, "services")
	})

	// Route for Courses Page
	router.GET("/courses", func(c *gin.Context) {
		renderPage(c, "courses")
	})

	// Generic route for other pages
	pages := []string{"exclusives", "photos", "support", "author"}
	for _, page := range pages {
		page := page // Fix closure issue
		router.GET("/"+page, func(c *gin.Context) {
			renderPage(c, page)
		})
	}

	// Custom 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found", "path": c.Request.URL.Path})
	})

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}

// Function to render pages with error checking
func renderPage(c *gin.Context, page string) {
	data, exists := pageData[page]
	if !exists {
		data = gin.H{"Title": "Page Not Found"}
		fmt.Println("WARNING: No data found for page:", page)
	}
	c.HTML(http.StatusOK, "layout.html", data)
}
