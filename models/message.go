package models

import (
	"site-portfolio/config"
)

type Message struct {
	ID      int
	Name    string
	Email   string
	Content string
}

func SaveMessage(name, email, message string) error {
	_, err := config.DB.Exec("INSERT INTO contact_messages (name, email, message) VALUES ($1, $2, $3)", name, email, message)
	return err
}
