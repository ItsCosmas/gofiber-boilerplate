package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	prod = "production"
)

// Config object
type Config struct {
	Env      string         `env:"ENV"`
	Pepper   string         `env:"PEPPER"`
	HMACKey  string         `env:"HMAC_KEY"`
	Postgres PostgresConfig `json:"postgres"`
	// Mailgun   MailgunConfig  `json:"mailgun"`
	JWTSecret string `env:"JWT_SIGN_KEY"`
	JWTIssuer string `env:"JWT_ISSUER"`
	Host      string `env:"APP_HOST"`
	Port      string `env:"APP_PORT"`
	// FromEmail string         `env:"EMAIL_FROM"`
}

// IsProd Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == prod
}

// LoadConfig gets config from .env
func LoadConfig() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	environmentPath := filepath.Join(currentPath, ".env")

	if err := godotenv.Load(environmentPath); err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:      os.Getenv("ENV"),
		Pepper:   os.Getenv("PEPPER"),
		HMACKey:  os.Getenv("HMAC_KEY"),
		Postgres: GetPostgresConfig(),
		// Mailgun:   GetMailgunConfig(),
		JWTSecret: os.Getenv("JWT_SIGN_KEY"),
		JWTIssuer: os.Getenv("JWT_ISSUER"),
		Host:      os.Getenv("APP_HOST"),
		Port:      os.Getenv("APP_PORT"),
		// FromEmail: os.Getenv("EMAIL_FROM"),
	}
}
