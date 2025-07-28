package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"net/smtp"
)

// SMTP настройки
var (
	smtpServer = "smtp.gmail.com"
	smtpPort   = "587"
	sender     = os.Getenv("SMTP_EMAIL")
	password   = os.Getenv("SMTP_PASSWORD")
	recipient  = os.Getenv("SMTP_RECIPIENT")
	// sender     = "peterhollander2025@gmail.com"
	// password   = "nxtf tdmd sxup eyta"
	// recipient  = "mehanic2000@gmail.com"

)

// Роуты для страницы контактов
func SetupContactRoutes(router *gin.Engine) {
	//router.GET("/contact", showContactForm)
	router.POST("/contact", handleFormSubmission)
}

// Отображение формы
func showContactForm(c *gin.Context) {
	tmpl, err := template.ParseFiles("templates/contact.html") // путь к HTML
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading form: %s", err)
		return
	}
	c.Writer.Header().Set("Content-Type", "text/html")
	tmpl.Execute(c.Writer, nil)
}

// Обработка формы
func handleFormSubmission(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")

	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	sender := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	recipient := os.Getenv("SMTP_RECIPIENT")

	subject := "New message from contact form"
	body := fmt.Sprintf("From: %s <%s>\n\n%s", name, email, message)
	msg := "Subject: " + subject + "\n\n" + body

	auth := smtp.PlainAuth("", sender, password, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, sender, []string{recipient}, []byte(msg))
	if err != nil {
		log.Println("SMTP Error:", err)
		c.String(http.StatusInternalServerError, "Failed to send email: %s", err)
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, "<h2>Thanks, %s!</h2><p>Your message has been sent.</p>", name)

	lang := getLanguage(c)
	fmt.Println("Selected language:", lang)

	// Получаем переводы для выбранного языка
	translationsForLang := translations[lang]

	// Динамически выбираем правильный шаблон для языка
	templateFile := fmt.Sprintf("templates/thankyou_%s.html", lang)

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Println("Template loading error:", err)
		c.String(http.StatusInternalServerError, "Error loading template: %s", err)
		return
	}

	// Передаем имя и язык в шаблон thankyou_<lang>.html
	err = tmpl.Execute(c.Writer, gin.H{
		"Name":          name,
		"Lang":          lang,
		"ThanksMessage": translationsForLang["thanks_message"],
		"MessageSent":   translationsForLang["message_sent"],
	})

	if err != nil {
		log.Println("Template execution error:", err)
		c.String(http.StatusInternalServerError, "Error rendering template: %s", err)
		return
	}

}

// Структура с переводами
var translations = map[string]map[string]string{
	"en": {
		"thanks_message": "Thanks !",
		"message_sent":   "Your message has been sent.",
	},
	"de": {
		"thanks_message": "Danke!",
		"message_sent":   "Ihre Nachricht wurde gesendet.",
	},
	"ar": {
		"thanks_message": "شكرًا!",
		"message_sent":   "تم إرسال رسالتك.",
	},
}

// func getLanguage(c *gin.Context) string {
// 	// Пробуем сначала из URL
// 	lang := c.Query("lang")
// 	if lang == "" {
// 		// Если нет в URL — пробуем из POST-данных формы
// 		lang = c.PostForm("lang")
// 	}

// 	log.Printf("Selected language: %s", lang)

// 	if lang != "de" && lang != "en" && lang != "ar" {
// 		lang = "en" // По умолчанию — английский
// 	}

// 	return lang
// }
