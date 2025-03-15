package config

import (
	"log"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
)

var (
	SMTPUserVault     string
	SMTPPasswordVault string
	DBConnStrVault    string
)

// LoadEnv loads secrets from either .env file or HashiCorp Vault
func LoadEnvVault() {
	// Load from .env file as fallback
	_ = godotenv.Load()

	// Load from Vault
	if err := loadFromVault(); err != nil {
		log.Printf("Vault error: %s. Falling back to .env", err)
	}

	// Load from environment variables as the final fallback
	SMTPUser = getEnv("SMTP_USER", SMTPUserVault)
	SMTPPassword = getEnv("SMTP_PASS", SMTPPasswordVault)
	DBConnStr = getEnv("DATABASE_URL", DBConnStrVault)
}

// loadFromVault fetches secrets from Vault
func loadFromVault() error {
	config := &api.Config{
		Address: os.Getenv("VAULT_ADDR"), // Example: "http://127.0.0.1:8200"
	}

	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	// Authenticate with Vault using a token
	client.SetToken(os.Getenv("VAULT_TOKEN")) // Or use Vault AppRole/Auth methods

	// Fetch the secrets from the secret engine (e.g., "secret/data/myapp")
	secret, err := client.Logical().Read("secret/data/myapp")
	if err != nil {
		return err
	}

	// Extract data
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return err
	}

	// Assign secrets to variables
	SMTPUser = data["SMTP_USER"].(string)
	SMTPPassword = data["SMTP_PASS"].(string)
	DBConnStr = data["DATABASE_URL"].(string)

	return nil
}

// getEnv is a helper to return the value from environment or fallback to the default
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
