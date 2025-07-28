package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func getLanguage(c *gin.Context) string {
// 	lang := c.Query("lang") // Get `lang` from URL query

// 	if lang != "de" && lang != "en" && lang != "nl"{
// 		lang = "en"
// 	}

// 	log.Printf("Selected language: %s", lang) // Log selected language
// 	return lang
// }

func renderPage(c *gin.Context, page string) {
	lang := getLanguage(c)                               // Get the selected language
	template := page + "_" + lang + ".html"              // Generate the template name
	log.Printf("Rendering template: %s", template)       // Log the template name being used
	c.HTML(http.StatusOK, template, gin.H{"Lang": lang}) // Render the template
}

func getLanguage(c *gin.Context) string {
	lang := c.Query("lang") // Get `lang` from URL query

	if lang == "" {
		// work with language template
		lang = c.PostForm("lang")
	}

	log.Printf("Selected language: %s", lang) // Log the selected language
	if lang != "de" && lang != "en" && lang != "ar" {
		lang = "en" // Default to English if no valid lang is found
	}
	return lang
}

// Existing Handlers Updated to Use renderPage

func HandleHome(c *gin.Context) {
	renderPage(c, "index")
}

func HandleContact(c *gin.Context) {
	renderPage(c, "contact")
}

func HandleServices(c *gin.Context) {
	renderPage(c, "services")
}

func HandleExclusives(c *gin.Context) {
	renderPage(c, "exclusives")
}

func HandleSupport(c *gin.Context) {
	renderPage(c, "support")
}

func HandleCourses(c *gin.Context) {
	renderPage(c, "courses")
}

func HandlePhotos(c *gin.Context) {
	renderPage(c, "photos")
}

func HandleVideos(c *gin.Context) {
	renderPage(c, "videos")
}

func HandleAuthor(c *gin.Context) {
	renderPage(c, "author")
}
