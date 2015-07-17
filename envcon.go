package envcon

import (
	"fmt"
	"os"
)

// Environment maps to dev, prod or testing
type Environment string

// Environment statuses
var (
	Development Environment = "development"
	Production  Environment = "production"
	Testing     Environment = "testing"
)

// Config house the environment variables used throughout the application
type Config struct {
	Environment      Environment
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseHost     string
	Port             string
	Key              string
	AdminEmail       string
	AppName          string
	AppURL           string
	EmailAddress     string
	EmailPassword    string
	EmailSMTP        string
	EmailPort        string
}

// Setup grabs the env variables and puts them in the config
func Setup(prefix string) Config {
	c := Config{}

	c.Environment = Environment(os.Getenv(fmt.Sprintf("%s_ENVIRONMENT", prefix)))
	c.DatabasePort = os.Getenv(fmt.Sprintf("%s_DATABASE_PORT", prefix))
	c.DatabaseUser = os.Getenv(fmt.Sprintf("%s_DATABASE_USER", prefix))
	c.DatabasePassword = os.Getenv(fmt.Sprintf("%s_DATABASE_PASSWORD", prefix))
	c.DatabaseName = os.Getenv(fmt.Sprintf("%s_DATABASE_NAME", prefix))
	c.DatabaseHost = os.Getenv(fmt.Sprintf("%s_DATABASE_HOST", prefix))
	c.Port = os.Getenv(fmt.Sprintf("%s_PORT", prefix))
	c.Key = os.Getenv(fmt.Sprintf("%s_KEY", prefix))
	c.AdminEmail = os.Getenv(fmt.Sprintf("%s_ADMIN_EMAIL", prefix))
	c.AppName = os.Getenv(fmt.Sprintf("%s_APP_NAME", prefix))
	c.AppURL = os.Getenv(fmt.Sprintf("%s_APP_URL", prefix))
	c.EmailAddress = os.Getenv(fmt.Sprintf("%s_EMAIL_ADDRESS", prefix))
	c.EmailPassword = os.Getenv(fmt.Sprintf("%s_EMAIL_PASSWORD", prefix))
	c.EmailSMTP = os.Getenv(fmt.Sprintf("%s_EMAIL_SMTP", prefix))
	c.EmailPort = os.Getenv(fmt.Sprintf("%s_EMAIL_PORT", prefix))

	return c
}
