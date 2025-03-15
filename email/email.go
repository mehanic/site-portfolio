package email

import (
	"fmt"
	"net/smtp"
	"site-portfolio/config"
)

// SendEmail sends an email notification
func SendEmail(name, senderEmail, message string) error {
	auth := smtp.PlainAuth("", config.SMTPUser, config.SMTPPassword, "smtp.gmail.com")
	to := []string{config.SMTPUser}
	subject := "New Contact Form Submission"
	body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage:\n%s", name, senderEmail, message)

	msg := []byte("Subject: " + subject + "\r\n" +
		"From: " + senderEmail + "\r\n" +
		"To: " + config.SMTPUser + "\r\n\r\n" +
		body)

	return smtp.SendMail("smtp.gmail.com:587", auth, senderEmail, to, msg)
}
